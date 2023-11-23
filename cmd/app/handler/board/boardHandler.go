package board

import (
	"cafe/internal/controller/board"
	"cafe/internal/controller/board/req"
	"cafe/internal/controller/board/res"
	"cafe/internal/controller/boardAction"
	"cafe/internal/controller/cafe"
	"cafe/internal/controller/member"
	"cafe/internal/controller/memberRole"
	"cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c          board.Controller
	cafeCon    cafe.Controller
	mCon       member.Controller
	mRoleCon   memberRole.Controller
	bActionCon boardAction.Controller
}

func NewHandler(c board.Controller, mCon member.Controller, cafeCon cafe.Controller, mRoleCon memberRole.Controller, bActionCon boardAction.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c, mCon: mCon, cafeCon: cafeCon, mRoleCon: mRoleCon, bActionCon: bActionCon}
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards/{id:[0-9]+}", h.getDetail).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards/{boardType:[0-9]+}", h.create).Methods(http.MethodPost)
	// 실제 작성자인지 확인할 로직이 필요
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/boards/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InvalidUserId        = "invalid user id"
	InvalidId            = "invalid board id"
	InvalidMember        = "invalid member"
	YouDonHavePermission = "You do not have permission"
	BoardNotFound        = "board not found"
	InvalidCafeId        = "invalid cafe id"
	InvalidBoardType     = "invalid board type"
	InternalServerError  = "internal server error"
)

// member 일경우 리스트는 조회가능 (제목 까지는 조회가능)

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InvalidUserId, http.StatusUnauthorized)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	mInfo, err := h.mCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	if mInfo.Id < 1 {
		http.Error(w, InvalidMember, http.StatusUnauthorized)
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

// cafe 주인이면 create
// cafe 주인이 아닐경우 boardAction 과 memberRole 과 비교해서 생성 가능 여부 설정
func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
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

// 글쓴이만 수정 가능
func (h Handler) patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InvalidUserId, http.StatusUnauthorized)
		return
	}
	mInfo, err := h.mCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mInfo.Id < 1 {
		http.Error(w, InvalidMember, http.StatusForbidden)
		return
	}

	detail, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if detail.Writer != mInfo.Id {
		http.Error(w, YouDonHavePermission, http.StatusForbidden)
		return
	}

	var d req.Patch
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Patch(r.Context(), id, mInfo.Id, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 글작성자이거나 카페 주인만 삭제 가능
func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InvalidUserId, http.StatusUnauthorized)
		return
	}
	ownerId, err := h.cafeCon.GetOwnerId(r.Context(), cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 카페 주인일경우 삭제 요청
	if userId == ownerId {
		err = h.c.Delete(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	}

	mInfo, err := h.mCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 작성자가 맞을경우 삭제 아닐경우 삭제 불가
	detail, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if detail.Writer != mInfo.Id {
		http.Error(w, YouDonHavePermission, http.StatusForbidden)
		return
	}
	err = h.c.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// 권한 확인필요
	//cafeId, err := strconv.Atoi(vars["cafeId"])
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, InvalidUserId, http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}

	dto, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if dto.Id < 1 {
		http.Error(w, BoardNotFound, http.StatusNotFound)
		return
	}
	mInfo, err := h.mCon.GetMemberInfo(r.Context(), cafeId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if mInfo.Id < 1 {
		http.Error(w, InvalidMember, http.StatusUnauthorized)
		return
	}
	roles, err := h.mRoleCon.GetOneMemberRoles(r.Context(), cafeId, mInfo.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	aInfo, err := h.bActionCon.GetInfo(r.Context(), cafeId, dto.BoardType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	readRoleIntArr := stringToIntArr(aInfo.ReadRoles)
	mRolesIntArr := stringToIntArr(roles.CafeRoleIds)
	readAble := checkContainIntArrToIntArr(mRolesIntArr, readRoleIntArr)
	if !readAble {
		http.Error(w, YouDonHavePermission, http.StatusForbidden)
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

func checkContainIntArrToIntArr(arr1 []int, arr2 []int) bool {
	for a1 := range arr1 {
		for a2 := range arr2 {
			if a1 == a2 {
				return true
			}
		}
	}
	return false
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
