package boardType

import (
	"cafe/internal/controller/boardType"
	"cafe/internal/controller/boardType/req"
	"cafe/internal/controller/cafe"
	"cafe/internal/controller/cafe/res"
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
	cafeCon cafe.Controller
}

func NewBoardTypeHandler(typeCon boardType.Controller, cafeCon cafe.Controller) http.Handler {
	r := mux.NewRouter()
	h := BoardTypeHandler{typeCon: typeCon, cafeCon: cafeCon}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types", h.getBoardList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types/{id:[0-9]+}", h.getDetail).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types", h.create).Methods(http.MethodPost)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types/{typeId:[0-9]+}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-types/{typeId:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InvalidCafeId          = "invalid cafe id"
	InvalidUserId          = "invalid user id"
	CafeNotFound           = "cafe not found"
	InternalServerError    = "internal server error"
	YouDoNotHavePermission = "You do not have permission"
	InvalidId              = "invalid board type id"
)

func (h BoardTypeHandler) getBoardList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}

	ok, err := h.cafeCon.IsExistsCafe(r.Context(), cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, CafeNotFound, http.StatusNotFound)
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
		http.Error(w, InternalServerError, http.StatusInternalServerError)
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
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InvalidUserId, http.StatusBadRequest)
		return
	}
	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, YouDoNotHavePermission, http.StatusForbidden)
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

func (h BoardTypeHandler) patch(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, YouDoNotHavePermission, http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	typeId, err := strconv.Atoi(vars["typeId"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}

	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, YouDoNotHavePermission, http.StatusForbidden)
		return
	}

	var pD req.PatchBoardTypeDto
	json.NewDecoder(r.Body).Decode(&pD)
	err = h.typeCon.Patch(r.Context(), cafeId, typeId, pD)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "no row") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h BoardTypeHandler) delete(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, YouDoNotHavePermission, http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	typeId, err := strconv.Atoi(vars["typeId"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}

	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, YouDoNotHavePermission, http.StatusForbidden)
		return
	}
	err = h.typeCon.Delete(r.Context(), cafeId, typeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h BoardTypeHandler) getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	dto, err := h.typeCon.GetDetail(r.Context(), cafeId, id)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(dto)
	if err != nil {
		log.Println("getDetail json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
