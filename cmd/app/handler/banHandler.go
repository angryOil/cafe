package handler

import (
	"cafe/internal/controller/ban"
	"cafe/internal/controller/ban/req"
	res2 "cafe/internal/controller/ban/res"
	"cafe/internal/controller/cafe"
	"cafe/internal/controller/cafe/res"
	"cafe/internal/controller/member"
	res3 "cafe/internal/controller/member/res"
	page2 "cafe/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BanHandler struct {
	cafeCon   cafe.Controller
	memberCon member.Controller
	banCon    ban.Controller
}

func NewBanHandler(banCon ban.Controller, memberCon member.Controller, cafeCon cafe.Controller) http.Handler {
	r := mux.NewRouter()
	h := BanHandler{
		cafeCon:   cafeCon,
		banCon:    banCon,
		memberCon: memberCon,
	}
	// 나의 카페 밴 확인
	r.HandleFunc("/cafes/ban/my", h.getBanListByUserId).Methods(http.MethodGet)
	// 밴하기 : 권한체크 cafeCon 에서 카페주인인지 확인 memberCon 에서 해당 userId_cafeId 유효한 멤버인지 확인 => 유효하면 밴처리
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/ban/admin", h.createBan).Methods(http.MethodPost)
	// 카페 벤 리스트 확인: 권한체크(카페 주인만)
	r.HandleFunc("/cafes/{cafeId:[0-9]+}/ban/admin", h.getCafeBanListByCafeId).Methods(http.MethodGet)
	return r
}

func (h BanHandler) createBan(w http.ResponseWriter, r *http.Request) {
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

	var bDto req.CreateBanDto
	err = json.NewDecoder(r.Body).Decode(&bDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 권한체크 (현재는 자기 카페인지) + 존재하는 카페인지 확인(같이됨)
	ok, err = h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "You do not have permission", http.StatusForbidden)
		return
	}

	// 존재하는 멤버인지 확인
	mDto, err := h.memberCon.GetInfoByCafeMemberId(r.Context(), cafeId, bDto.MemberId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if mDto.Id == 0 {
		http.Error(w, "member is not exists", http.StatusNotFound)
		return
	}

	// 저장
	err = h.banCon.CreateBan(r.Context(), mDto.UserId, cafeId, bDto)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("createBan err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// todo 리펙토링 하기
func (h BanHandler) getBanListByUserId(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	reqPage := page2.GetPageReqByRequest(r)
	banListDtos, count, err := h.banCon.GetMyBanListAndCount(r.Context(), userId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 탐색용 cafeIds arr
	cafeIds := make([]int, len(banListDtos))
	for i, d := range banListDtos {
		cafeIds[i] = d.CafeId
	}

	cafeNameDtos, err := h.cafeCon.GetCafesByCafeIds(r.Context(), cafeIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cafeNameMap := make(map[int]string, len(cafeNameDtos))
	for _, cN := range cafeNameDtos {
		cafeNameMap[cN.Id] = cN.Name
	}

	banListDtoMap := make(map[int]res2.BanListDto, len(banListDtos))
	for _, d := range banListDtos {
		banListDtoMap[d.CafeId] = d
	}

	detailListMap := make(map[int]res2.BanDetailDto, len(banListDtos))

	for i, b := range banListDtoMap {
		name, ok := cafeNameMap[i]
		if !ok {
			name = ""
		}
		detailListMap[i] = b.ToDetailDto(name)
	}
	detailList := make([]res2.BanDetailDto, 0)
	for _, m := range detailListMap {
		detailList = append(detailList, m)
	}
	listCountDto := res.NewListTotalDto(detailList, count)
	data, err := json.Marshal(listCountDto)

	if err != nil {
		log.Println("getBanListByUserId json.Marshal err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h BanHandler) getCafeBanListByCafeId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	isMine, err := h.cafeCon.CheckIsMine(r.Context(), userId, cafeId)
	if err != nil {
		log.Println("getCafeBanListByCafeId CheckIsMine err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if !isMine {
		http.Error(w, "You do not have permission", http.StatusForbidden)
		return
	}
	reqPage := page2.GetPageReqByRequest(r)

	banAdminListDtos, total, err := h.banCon.GetBanListByCafeId(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(banAdminListDtos) == 0 {
		data, err := json.Marshal(res.NewListTotalDto([]res2.BanAdminDetailDto{}, total))
		if err != nil {
			log.Println("getCafeBanListByCafeId json.Marshal err: ", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
		return
	}
	memberIds := make([]int, len(banAdminListDtos))
	for i, b := range banAdminListDtos {
		memberIds[i] = b.MemberId
	}

	memberDtos, err := h.memberCon.GetMembersByMemberIds(r.Context(), memberIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	memberDtosMap := make(map[int]res3.MemberInfoDto, len(memberDtos))
	for _, m := range memberDtos {
		memberDtosMap[m.Id] = m
	}

	detailMap := make(map[int]res2.BanAdminDetailDto, len(banAdminListDtos))
	for _, b := range banAdminListDtos {
		dto, ok := memberDtosMap[b.MemberId]
		name := ""
		if ok {
			name = dto.Nickname
		}
		detailMap[b.MemberId] = b.ToDetailDto(name)
	}

	detailList := make([]res2.BanAdminDetailDto, 0, len(detailMap))
	for _, d := range detailMap {
		detailList = append(detailList, d)
	}

	totalList := res.NewListTotalDto(detailList, total)

	data, err := json.Marshal(totalList)
	if err != nil {
		log.Println("getCafeBanListByCafeId err: ", err)
		http.Error(w, "server internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
