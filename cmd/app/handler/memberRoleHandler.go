package handler

import (
	"cafe/internal/cli/memberRole/cDto"
	"cafe/internal/controller/cafe"
	"cafe/internal/controller/cafeRole"
	"cafe/internal/controller/cafeRole/res"
	"cafe/internal/controller/member"
	"cafe/internal/controller/memberRole"
	req2 "cafe/internal/controller/memberRole/req"
	res2 "cafe/internal/controller/memberRole/res"
	"cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MemberRoleHandler struct {
	cafeCon  cafe.Controller //
	memCon   member.Controller
	cRoleCon cafeRole.Controller
	mRoleCon memberRole.Controller
}

func NewMemberRoleHandler(cafeCon cafe.Controller, memCon member.Controller, cRoleCon cafeRole.Controller, mRoleCon memberRole.Controller) http.Handler {
	r := mux.NewRouter()
	h := MemberRoleHandler{cafeCon: cafeCon, memCon: memCon, mRoleCon: mRoleCon}
	// 카페 전체인원관련 권한 확인
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/member-roles", h.getMembersRoles).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/member-roles/{memberId:[0-9]+}", h.getOneMemberRoles).Methods(http.MethodGet)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/member-roles/{memberId:[0-9]+}/{mRoleId:[0-9]+}", h.delete).Methods(http.MethodDelete)
	return r
}

func (h MemberRoleHandler) getOneMemberRoles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, "invalid member id", http.StatusBadRequest)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id ", http.StatusBadRequest)
		return
	}

	d, err := h.mRoleCon.GetOneMemberRoles(r.Context(), cafeId, memberId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	memberIdsArr := stringToIntArr(d.CafeRoleIds)

	reqPage := page.GetPageReqByRequest(r)
	cafeRoles, _, err := h.cRoleCon.GetList(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	roleMap := make(map[int]res.RoleDto)
	for _, cRole := range cafeRoles {
		roleMap[cRole.Id] = cRole
	}

	roleArr := make([]res2.Role, 0)
	for mId := range memberIdsArr {
		role, ok := roleMap[mId]
		if !ok {
			continue
		}
		roleArr = append(roleArr, res2.Role{
			RoleId: role.Id,
			Name:   role.Name,
		})
	}

	memberRoleArrDto := d.ToRoleArrDto(roleArr)

	data, err := json.Marshal(memberRoleArrDto)
	if err != nil {
		log.Println("getOneMemberRoles json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// todo 다른 방법있을지 고민하기  심각하게 성능이 걱정되는 메소드
func (h MemberRoleHandler) getMembersRoles(w http.ResponseWriter, r *http.Request) {
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
	// 카페 주인인지 확인
	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "you do not have permission", http.StatusForbidden)
		return
	}

	reqPage := page.GetPageReqByRequest(r)
	mRoleDetailDtos, memberTotalCount, err := h.mRoleCon.GetRolesByCafeId(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cafeRoleDtos, _, err := h.cRoleCon.GetList(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// todo 아래부턴 리팩토링 필요
	cafeRoleDtoMap := make(map[int]res.RoleDto)
	for _, r := range cafeRoleDtos {
		cafeRoleDtoMap[r.Id] = r
	}

	mRoleArrDto := make([]res2.DetailRoleArrDto, 0)
	for _, m := range mRoleDetailDtos {
		roleDtos := make([]res2.Role, 0)
		intArr := stringToIntArr(m.CafeRoleIds)
		for i := range intArr {
			cRole, ok := cafeRoleDtoMap[i]
			if !ok {
				continue
			}
			roleDtos = append(roleDtos, res2.Role{
				RoleId: cRole.Id,
				Name:   cRole.Name,
			})
		}
		mRoleArrDto = append(mRoleArrDto, m.ToRoleArrDto(roleDtos))
	}

	data, err := json.Marshal(cDto.NewListTotalDto(mRoleArrDto, memberTotalCount))
	if err != nil {
		log.Println("getMembersRoles json.Marshal err: ", err)
		http.Error(w, "server internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func stringToIntArr(s string) []int {
	str := strings.ReplaceAll(s, " ", "")
	arr := strings.Split(str, ",")
	intArr := make([]int, 0)
	for _, s := range arr {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		intArr = append(intArr, i)
	}
	return intArr
}

func (h MemberRoleHandler) upsert(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "you do not have permission", http.StatusForbidden)
		return
	}

	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, "invalid member id", http.StatusBadRequest)
		return
	}

	ok, err = h.memCon.CheckExistsMemberByMemberId(r.Context(), cafeId, memberId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "member not found", http.StatusNotFound)
		return
	}

	var d req2.PutMemberRoleDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.mRoleCon.PutRole(r.Context(), cafeId, memberId, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h MemberRoleHandler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	mRoleId, err := strconv.Atoi(vars["mRoleId"])
	if err != nil {
		http.Error(w, "invalid member role id", http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, "invalid member id", http.StatusBadRequest)
		return
	}
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "you do not permission", http.StatusForbidden)
		return
	}

	err = h.mRoleCon.Delete(r.Context(), cafeId, memberId, mRoleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
