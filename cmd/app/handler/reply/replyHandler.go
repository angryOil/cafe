package reply

import (
	"cafe/internal/controller/reply"
	"cafe/internal/controller/reply/res"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	c reply.Controller
}

func NewHandler(c reply.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/replies/{boardId:[0-9]+}", h.getList).Methods(http.MethodGet)
	return r
}

const (
	InvalidCafeId       = "invalid cafe id"
	InvalidBoardId      = "invalid board id"
	InternalServerError = "internal server error"
)

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//cafeId, err := strconv.Atoi(vars["cafeId"])
	//if err != nil {
	//	http.Error(w, InvalidCafeId, http.StatusBadRequest)
	//	return
	//}
	boardId, err := strconv.Atoi(vars["boardId"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}

	list, total, err := h.c.GetList(r.Context(), boardId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listTotalDto := res.ListTotalDto{
		Content: list,
		Total:   total,
	}
	data, err := json.Marshal(listTotalDto)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
