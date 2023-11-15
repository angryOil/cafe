package boardAction

import (
	"cafe/internal/controller/boardAction"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// 2가지 역활을 합니다
// 1. 보드 타입에 따른 crud 권한 저장
// 2. 권한 비교결과 알려주는 기능 (단 작성자는 무조건 read,delete 가능)

type BoardActionHandler struct {
	c boardAction.Controller
}

func NewBoardActionHandler(c boardAction.Controller) http.Handler {
	r := mux.NewRouter()
	h := BoardActionHandler{c: c}

	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}", h.getInfo).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}", h.create).Methods(http.MethodPost)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InternalServerError = "internal server error"
)

func (h BoardActionHandler) getInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	boardTypeId, err := strconv.Atoi(vars["typeId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	d, err := h.c.GetInfo(r.Context(), cafeId, boardTypeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(&d)
	if err != nil {
		log.Println("getInfo json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h BoardActionHandler) create(w http.ResponseWriter, r *http.Request) {

}

func (h BoardActionHandler) patch(w http.ResponseWriter, r *http.Request) {

}

func (h BoardActionHandler) delete(w http.ResponseWriter, r *http.Request) {

}
