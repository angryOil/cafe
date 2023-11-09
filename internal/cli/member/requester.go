package member

import (
	"bytes"
	dto2 "cafe/internal/cli/member/dto"
	"cafe/internal/cli/member/req"
	"cafe/internal/cli/member/res"
	"cafe/internal/domain"
	"cafe/internal/page"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// memberURL
var memberURL = "http://localhost:8084/members"
var adminMemberUrl = "http://localhost:8084/admin/members"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

const (
	InternalServerError = "internal server error"
)

// 해당 카페의 member(자기정보) 정보조회

func (r Requester) GetCafeMyInfo(ctx context.Context, cafeId, userId int) (res.GetCafeMyInfo, error) {
	reqUrl := fmt.Sprintf("%s/%d/info/%d", memberURL, cafeId, userId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("cli GetCafeMyInfo NewRequest err: ", err)
		return res.GetCafeMyInfo{}, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("cli GetCafeMyInfo NewRequest err: ", err)
		return res.GetCafeMyInfo{}, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res.GetCafeMyInfo{}, errors.New(InternalServerError)
	}

	var md res.GetCafeMyInfo
	err = json.NewDecoder(resp.Body).Decode(&md)
	if err != nil {
		log.Println("GetCafeMyInfo json decode err: ", err)
		return res.GetCafeMyInfo{}, errors.New(InternalServerError)
	}
	return md, nil
}

func (r Requester) GetCafeIdsAndTotalByUserId(ctx context.Context, userId int, reqPage page.ReqPage) ([]int, int, error) {
	reqUrl := fmt.Sprintf("%s/list/%d?page=%d&size=%d", memberURL, userId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId NewRequest err: ", err)
		return []int{}, 0, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId DefaultClient.Do err: ", err)
		return []int{}, 0, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	var iTD domain.IdsTotalDomain
	err = json.NewDecoder(resp.Body).Decode(&iTD)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId json.decode err: ", err)
		return []int{}, 0, errors.New(InternalServerError)
	}
	return iTD.Ids, iTD.Total, nil
}

func (r Requester) JoinCafe(ctx context.Context, d req.JoinCafe) error {
	reqUrl := fmt.Sprintf("%s/%d/join/%d", memberURL, d.CafeId, d.UserId)
	jd := dto2.ToJoinMemberDto(d)
	data, err := json.Marshal(jd)
	if err != nil {
		log.Println("JoinCafe json.Marshal err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("POST", reqUrl, bytes.NewReader(data))
	if err != nil {
		log.Println("JoinCafe NewRequest err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("JoinCafe DefaultClient.Do err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("JoinCafe readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) GetCafeMemberListCount(ctx context.Context, cafeId int, isBanned bool, reqPage page.ReqPage) (res.MemberInfoListCountDto, error) {
	reqUrl := fmt.Sprintf("%s/admin/%d?ban=%t&page=%d&size=%d", memberURL, cafeId, isBanned, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetCafeMemberListCount NewRequest err: ", err)
		return res.MemberInfoListCountDto{}, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetCafeMemberListCount defaultClient do err: ", err)
		return res.MemberInfoListCountDto{}, errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetCafeMemberListCount readBody err: ", err)
			return res.MemberInfoListCountDto{}, errors.New(InternalServerError)
		}
		return res.MemberInfoListCountDto{}, errors.New(string(readBody))
	}
	var listCount res.MemberInfoListCountDto
	err = json.NewDecoder(resp.Body).Decode(&listCount)
	if err != nil {
		log.Println("GetCafeMemberListCount json decode err: ", err)
	}
	return listCount, nil
}

func (r Requester) PatchMember(ctx context.Context, d req.PatchMember) error {
	reqUrl := fmt.Sprintf("%s/%d", memberURL, d.MemberId)
	dto := d.ToPatchDto()
	data, err := json.Marshal(dto)
	if err != nil {
		log.Println("PatchMember json marshal err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("PATCH", reqUrl, bytes.NewReader(data))
	if err != nil {
		log.Println("PatchMember NewRequest marshal err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("PatchMember DefaultClient marshal err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("PatchMember readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) GetMemberByCafeMemberId(ctx context.Context, cafeId int, memberId int) (res.GetMemberByCafeMemberId, error) {
	reqUrl := fmt.Sprintf("%s/%d", memberURL, memberId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetMemberByCafeMemberId NewRequest err: ", err)
		return res.GetMemberByCafeMemberId{}, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetMemberByCafeMemberId DefaultClient.Do err: ", err)
		return res.GetMemberByCafeMemberId{}, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetMemberByCafeMemberId readBody err: ", err)
			return res.GetMemberByCafeMemberId{}, errors.New(InternalServerError)
		}
		return res.GetMemberByCafeMemberId{}, errors.New(string(readBody))
	}

	var dto res.GetMemberByCafeMemberId
	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		log.Println("GetMemberByCafeMemberId json.NewDecoder err: ", err)
		return res.GetMemberByCafeMemberId{}, errors.New(InternalServerError)
	}
	return dto, nil
}

func (r Requester) GetMemberListByMemberIds(ctx context.Context, ids []int) ([]res.GetMemberListByMemberIds, error) {
	idsStr := arrayToString(ids, ",")
	reqUrl := fmt.Sprintf("%s/admin?memberIds=%s", memberURL, idsStr)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetMemberListByMemberIds NewRequest err: ", err)
		return []res.GetMemberListByMemberIds{}, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetMemberListByMemberIds DefaultClient err: ", err)
		return []res.GetMemberListByMemberIds{}, errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetMemberListByMemberIds readAll err: ", err)
			return []res.GetMemberListByMemberIds{}, errors.New(InternalServerError)
		}
		return []res.GetMemberListByMemberIds{}, errors.New(string(readBody))
	}
	var list []res.GetMemberListByMemberIds
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		log.Println("GetMemberListByMemberIds decode err: ", err)
		return []res.GetMemberListByMemberIds{}, errors.New(InternalServerError)
	}
	return list, nil
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
