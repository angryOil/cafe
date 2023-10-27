package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/member"
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
	c      member.Controller
	cafeCo controller.CafeController
}

func NewMemberHandler(c member.Controller, cc controller.CafeController) http.Handler {
	m := mux.NewRouter()
	h := MemberHandler{c: c, cafeCo: cc}
	//  나의 카페 리스트 조회 (cafe_ids)
	m.HandleFunc("/cafes/members/my", h.getMyCafeList).Methods(http.MethodGet)
	// 해당카페 내 정보 조회
	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/info", h.getMemberInfo).Methods(http.MethodGet)
	// 카페 가입 신청
	//m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/join", h.getMemberInfo).Methods(http.MethodPost)

	// 관리자
	// 카페 가입 멤버리스트 조회
	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/admin", h.getMemberList).Methods(http.MethodGet)
	m.HandleFunc("/cafes/{cafeId:[0-9]+}/members/admin", h.patchMember).Methods(http.MethodPatch)

	return m
}

func (h MemberHandler) getMyCafeList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 0
	}
	idsTotalDto, err := h.c.GetMyCafeIds(r.Context(), userId, page2.NewReqPage(page, size))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cafeDtos, err := h.cafeCo.GetCafesByCafeIds(r.Context(), idsTotalDto.Ids)
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
	log.Println("member info")
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
	log.Println("cafeId:", cafeId, "userId", userId)
	dto, err := h.c.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("h.c.getMemberInfo err: ", err)
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

func (h MemberHandler) getMemberList(w http.ResponseWriter, r *http.Request) {

}

func (h MemberHandler) patchMember(w http.ResponseWriter, r *http.Request) {

}
