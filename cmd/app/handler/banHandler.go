package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/ban"
	"cafe/internal/controller/ban/req"
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

type BanHandler struct {
	cafeCon   controller.CafeController
	memberCon member.Controller
	banCon    ban.Controller
}

func NewBanHandler(banCon ban.Controller, memberCon member.Controller, cafeCon controller.CafeController) http.Handler {
	r := mux.NewRouter()
	h := BanHandler{
		cafeCon:   cafeCon,
		banCon:    banCon,
		memberCon: memberCon,
	}
	// 나의 카페 밴 확인
	r.HandleFunc("/cafes/ban/my", h.getBanListByUserId).Methods(http.MethodGet)
	// 밴하기 : 권한체크 cafeCon 에서 카페주인인지 확인 memberCon 에서 해당 userId_cafeId 유효한 멤버인지 확인 => 유효하면 밴처리
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/ban/admin", h.createBan).Methods(http.MethodPost)
	// 카페 벤 리스트 확인: 권한체크(카페 주인만)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/ban/admin", h.getCafeBanListByCafeId).Methods(http.MethodGet)
	return r
}

func (h BanHandler) createBan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var bDto req.CreateBanDto
	err = json.NewDecoder(r.Body).Decode(&bDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 권한체크 (현재는 자기 카페인지) + 존재하는 카페인지 확인(같이됨)
	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "You do not have permission", http.StatusForbidden)
		return
	}

	// 존재하는 멤버인지 확인
	mDto, err := h.memberCon.GetInfoByCafeMemberId(r.Context(), cafeId, bDto.MemberId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if mDto.Id == 0 {
		http.Error(w, "member is not exists", http.StatusNotFound)
		return
	}

	// 저장
	err = h.banCon.CreateBan(r.Context(), mDto.UserId, cafeId, bDto)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("createBan err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h BanHandler) getBanListByUserId(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	reqPage := page2.GetPageReqByRequest(r)
	list, count, err := h.banCon.GetMyBanListAndCount(r.Context(), userId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	listTotalDto := res.NewListTotalDto(list, count)
	data, err := json.Marshal(listTotalDto)
	if err != nil {
		log.Println("getBanListByUserId json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h BanHandler) getCafeBanListByCafeId(writer http.ResponseWriter, request *http.Request) {

}
