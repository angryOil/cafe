package reply

import (
	"cafe/internal/controller/reply"
	"cafe/internal/controller/reply/res"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c reply.Controller
}

func NewHandler(c reply.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/replies/{boardId:[0-9]+}", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/replies/cnt", h.getCount).Methods(http.MethodGet)
	return r
}

const (
	InvalidCafeId       = "invalid cafe id"
	InvalidBoardId      = "invalid board id"
	InvalidBoardIds     = "invalid board ids"
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

func (h Handler) getCount(w http.ResponseWriter, r *http.Request) {
	boardIds := r.URL.Query().Get("board-ids")
	if boardIds == "" {
		http.Error(w, InvalidBoardIds, http.StatusBadRequest)
		return
	}

	arr := stringToIntArr(boardIds)
	if len(arr) == 0 {
		return
	}

	list, err := h.c.GetCount(r.Context(), arr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	countListDto := res.CountListDto{Content: list}
	data, err := json.Marshal(countListDto)
	if err != nil {
		log.Println("getCount json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func stringToIntArr(s string) []int {
	s = strings.ReplaceAll(s, " ", "")
	sArr := strings.Split(s, ",")
	intArr := make([]int, 0)
	for _, s := range sArr {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		intArr = append(intArr, i)
	}
	return intArr
}
