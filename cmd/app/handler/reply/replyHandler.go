package reply

import (
	"cafe/internal/controller/member"
	"cafe/internal/controller/reply"
	"cafe/internal/controller/reply/req"
	"cafe/internal/controller/reply/res"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c    reply.Controller
	mCon member.Controller
}

func NewHandler(mCon member.Controller, c reply.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c, mCon: mCon}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/replies/{boardId:[0-9]+}", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/replies/cnt", h.getCount).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/replies/{boardId:[0-9]+}", h.create).Methods(http.MethodPost)
	return r
}

const (
	InvalidCafeId          = "invalid cafe id"
	YouDoNotHavePermission = "You do not have permission"
	InvalidUserId          = "invalid user id"
	InvalidBoardId         = "invalid board id"
	InvalidBoardIds        = "invalid board ids"
	InternalServerError    = "internal server error"
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

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InvalidUserId, http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardId, err := strconv.Atoi(vars["boardId"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}
	mInfo, err := h.mCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if mInfo.Id < 1 {
		http.Error(w, YouDoNotHavePermission, http.StatusForbidden)
		return
	}
	var c req.Create
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Create(r.Context(), cafeId, boardId, mInfo.Id, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
