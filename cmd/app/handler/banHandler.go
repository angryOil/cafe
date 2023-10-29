package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/ban"
	"cafe/internal/controller/ban/req"
	"cafe/internal/controller/member"
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

// 밴
// 1.밴하기 : 권한체크 cafeCon 에서 카페주인인지 확인 memberCon 에서 해당 userId_cafeId 유효한 멤버인지 확인 => 유효하면 밴처리

func NewBanHandler(banCon ban.Controller, memberCon member.Controller, cafeCon controller.CafeController) http.Handler {
	r := mux.NewRouter()
	h := BanHandler{
		cafeCon:   cafeCon,
		banCon:    banCon,
		memberCon: memberCon,
	}
	// 밴하기
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/ban/admin", h.createBan).Methods(http.MethodPost)
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
