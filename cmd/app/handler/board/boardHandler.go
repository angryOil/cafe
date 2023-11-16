package board

import (
	"cafe/internal/controller/board"
	"cafe/internal/controller/board/req"
	"cafe/internal/controller/board/res"
	"cafe/internal/controller/member"
	"cafe/internal/domain/ban"
	"cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	c    board.Controller
	mCon member.Controller
}

func NewHandler(c board.Controller, mCon member.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c, mCon: mCon}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards/{boardType:[0-9]+}", h.create).Methods(http.MethodPost)
	return r
}

const (
	InvalidUserId       = "invalid user id"
	InvalidBoardType    = "invalid board type"
	InternalServerError = "internal server error"
)

// 관리자 인가? 관리자일 경우 모든 메소드 허용 cafeCon
// 회원인가? 회원이 아닐경우 거부 memberCon
// 권한이 있는가? BoardAction + memberRole 을 비교해서 확인

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqPage := page.GetPageReqByRequest(r)
	q := r.URL.Query()
	boardType, err := strconv.Atoi(q.Get("board-type"))
	if err != nil {
		boardType = 0
	}
	writer, err := strconv.Atoi(q.Get("writer"))
	if err != nil {
		writer = 0
	}
	list, total, err := h.c.GetList(r.Context(), cafeId, boardType, writer, reqPage)
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
		log.Println("getList json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, ban.InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardType, err := strconv.Atoi(vars["boardType"])
	if err != nil {
		http.Error(w, InvalidBoardType, http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	mInfo, err := h.mCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var c req.Create
	err = json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Create(r.Context(), cafeId, boardType, mInfo.Id, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}