package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/req"
	page2 "cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CafeHandler struct {
	c controller.CafeController
}

func NewCafeHandler(c controller.CafeController) http.Handler {
	m := mux.NewRouter()
	h := CafeHandler{c: c}
	// 전체 카페를 조회
	m.HandleFunc("/cafes", h.getList).Methods(http.MethodGet)
	// 카페 생성
	m.HandleFunc("/cafes", h.createCafe).Methods(http.MethodPost)
	m.HandleFunc("/cafes/{id:[0-9]+}", h.getDetail).Methods(http.MethodGet)
	m.HandleFunc("/cafes/{id:[0-9]+}", h.updateCafe).Methods(http.MethodPut)
	// todo cafe 삭제 기능은 좀더 생각 해보기
	//m.HandleFunc("/cafes/{id:[0-9]+}", h.deleteCafe).Methods(http.MethodDelete)
	return m
}

func (h CafeHandler) createCafe(w http.ResponseWriter, r *http.Request) {
	reqDto := req.CreateCafeDto{}
	err := json.NewDecoder(r.Body).Decode(&reqDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.c.CreateCafe(r.Context(), reqDto)
	if err != nil {
		if strings.Contains(err.Error(), "internal server error") {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if strings.Contains(err.Error(), "duplicate key") {
			http.Error(w, "이미 존재하는 카페 이름입니다.", http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// 전체 카페리스트
func (h CafeHandler) getList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 0
	}
	reqPage := page2.NewReqPage(page, size)

	cafeList, count, err := h.c.GetCafes(r.Context(), reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageNation := page2.GetPagination(cafeList, reqPage, count)
	data, err := json.Marshal(pageNation)
	if err != nil {
		log.Println("getList marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// todo 삭제 보단 비활성화를 시킨후 일정기간후에 삭제 되게 만들기
func (h CafeHandler) deleteCafe(w http.ResponseWriter, r *http.Request) {

}

func (h CafeHandler) getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid cafe id ", http.StatusBadRequest)
		return
	}
	detail, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "id is zero") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(detail)
	if err != nil {
		log.Println("getDetail json marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h CafeHandler) updateCafe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}
	var dto req.UpdateCafeDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Update(r.Context(), dto, id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid user") || strings.Contains(err.Error(), "empty") || strings.Contains(err.Error(), "zero") || strings.Contains(err.Error(), "not exists") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
