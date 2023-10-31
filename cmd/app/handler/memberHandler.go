package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/member"
	"cafe/internal/controller/member/req"
	"cafe/internal/controller/res"
	page2 "cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MemberHandler struct {
	memberCon member.Controller
	cafeCon   controller.CafeController
}

func NewMemberHandler(c member.Controller, cc controller.CafeController) http.Handler {
	m := mux.NewRouter()
	h := MemberHandler{memberCon: c, cafeCon: cc}
	//  나의 카페 리스트 조회 (cafe_ids)
	m.HandleFunc("/cafes/members/my", h.getMyCafeList).Methods(http.MethodGet)
	// 해당카페 내 정보 조회
	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/info", h.getMemberInfo).Methods(http.MethodGet)
	// 카페 가입 신청
	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/join", h.joinCafe).Methods(http.MethodPost)

	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members", h.patchMember).Methods(http.MethodPatch)

	// 관리자
	// 카페 가입 멤버리스트 조회 현재로선 cafe 주인인지 확인 하고 요청
	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/admin", h.getMemberList).Methods(http.MethodGet)
	//m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/admin", h.adminPatchMember).Methods(http.MethodPatch)

	return m
}

func (h MemberHandler) getMyCafeList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	reqPage := page2.GetPageReqByRequest(r)
	idsTotalDto, err := h.memberCon.GetMyCafeIds(r.Context(), userId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cafeDtos, err := h.cafeCon.GetCafesByCafeIds(r.Context(), idsTotalDto.Ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listTotalDto := res.NewListTotalDto(cafeDtos, idsTotalDto.Total)
	data, err := json.Marshal(listTotalDto)
	if err != nil {
		log.Println("getMyCafeList json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h MemberHandler) getMemberInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id ", http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	dto, err := h.memberCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("h.memberCon.getMemberInfo err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(&dto)
	if err != nil {
		log.Println("getMemberInfo : ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h MemberHandler) joinCafe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id ", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id ", http.StatusBadRequest)
		return
	}

	ok, err = h.cafeCon.IsExistsCafe(r.Context(), cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "cafe not exists", http.StatusNotFound)
		return
	}

	var dto req.JoinMemberDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Println("joinCafe json.NewDecoder err: ", err)
		http.Error(w, "server internal error", http.StatusInternalServerError)
		return
	}

	if dto.Nickname == "" {
		http.Error(w, "nickname is empty", http.StatusBadRequest)
		return
	}

	err = h.memberCon.JoinCafe(r.Context(), userId, cafeId, dto)
	if err != nil {
		if strings.Contains(err.Error(), "empty") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if strings.Contains(err.Error(), "json ") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// 아래부턴 admin 기능

func (h MemberHandler) getMemberList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id ", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}

	isMine, err := h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isMine {
		http.Error(w, "You do not have permission", http.StatusForbidden)
		return
	}
	// 아래부터는  cafe api 에서 권한 체크를 했다고 산정하고 실행됨
	isBanned := "true" == r.URL.Query().Get("ban")
	reqPage := page2.GetPageReqByRequest(r)
	mInfoDtos, count, err := h.memberCon.GetCafeMemberListCount(r.Context(), cafeId, isBanned, reqPage)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalDto := res.NewListTotalDto(mInfoDtos, count)
	data, err := json.Marshal(&totalDto)
	if err != nil {
		log.Println("getMemberList json.Marshal err :", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h MemberHandler) patchMember(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id ", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}

	isMine, err := h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isMine {
		http.Error(w, "You do not have permission", http.StatusForbidden)
		return
	}

	var dto req.PatchMemberDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.memberCon.PatchMember(r.Context(), userId, cafeId, dto)
	if err != nil {
		if strings.Contains(err.Error(), "no row") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
