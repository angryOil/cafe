package handler

import (
	"cafe/internal/controller"
	"cafe/internal/controller/boardType"
	"cafe/internal/controller/boardType/req"
	"cafe/internal/controller/res"
	"cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BoardTypeHandler struct {
	typeCon boardType.Controller
	cafeCon controller.CafeController
}

func NewBoardTypeHandler(typeCon boardType.Controller, cafeCon controller.CafeController) http.Handler {
	r := mux.NewRouter()
	h := BoardTypeHandler{typeCon: typeCon, cafeCon: cafeCon}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types", h.getBoardList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types", h.create).Methods(http.MethodPost)
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

func (h BoardTypeHandler) create(w http.ResponseWriter, r *http.Request) {
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
	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "You do not have permission", http.StatusForbidden)
		return
	}

	ownerId, err := h.cafeCon.GetOwnerId(r.Context(), cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var d req.CreateBoardTypeDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println("create json.NewDecoder err: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.typeCon.Create(r.Context(), cafeId, ownerId, d)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
