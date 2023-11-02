package handler

import (
	"cafe/internal/cli/memberRole/cDto"
	"cafe/internal/controller"
	"cafe/internal/controller/cafeRole"
	"cafe/internal/controller/cafeRole/res"
	"cafe/internal/controller/member"
	"cafe/internal/controller/memberRole"
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
	cafeCon  controller.CafeController //
	memCon   member.Controller
	cRoleCon cafeRole.Controller
	mRoleCon memberRole.Controller
}

func NewMemberRoleHandler(cafeCon controller.CafeController, memCon member.Controller, cRoleCon cafeRole.Controller, mRoleCon memberRole.Controller) http.Handler {
	r := mux.NewRouter()
	h := MemberRoleHandler{cafeCon: cafeCon, memCon: memCon, mRoleCon: mRoleCon}
	// 카페 전체인원관련 권한 확인
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/member-roles", h.getMembersRoles).Methods(http.MethodGet)
	return r
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
	cafeRoleDtoMap := make(map[int]res.RoleDto)
	for _, r := range cafeRoleDtos {
		cafeRoleDtoMap[r.Id] = r
	}

	mRoleArrDto := make([]res2.DetailRoleArrDto, 0)
	for _, m := range mRoleDetailDtos {
		roleDtos := make([]res2.Role, 0)
		intArr := stringToIntArr(m.CafeRoleIds)
		for i := range intArr {
			cafeRole, ok := cafeRoleDtoMap[i]
			if !ok {
				continue
			}
			roleDtos = append(roleDtos, res2.Role{
				RoleId: cafeRole.Id,
				Name:   cafeRole.Name,
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
