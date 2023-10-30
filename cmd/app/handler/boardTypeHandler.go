package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/boardType"
	"cafe/internal/controller/res"
	"cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type BoardTypeHandler struct {
	typeCon boardType.Controller
	cafeCon controller.CafeController
}

func NewBoardTypeHandler(typeCon boardType.Controller, cafeCon controller.CafeController) http.Handler {
	r := mux.NewRouter()
	h := BoardTypeHandler{typeCon: typeCon, cafeCon: cafeCon}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types", h.getBoardList).Methods(http.MethodGet)
	return r
}

func (h BoardTypeHandler) getBoardList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}

	ok, err := h.cafeCon.IsExistsCafe(r.Context(), cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "cafe not found", http.StatusNotFound)
		return
	}

	reqPage := page.GetPageReqByRequest(r)

	dtos, total, err := h.typeCon.GetList(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(res.NewListTotalDto(dtos, total))
	if err != nil {
		log.Println("getBoardList json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
