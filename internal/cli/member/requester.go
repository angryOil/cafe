package member

import (
	"bytes"
	dto2 "cafe/internal/cli/member/dto"
	"cafe/internal/domain"
	"cafe/internal/page"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// memberURL
var memberURL = "http://localhost:8084/members"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

// 해당 카페의 member(자기정보) 정보조회

func (r Requester) GetCafeMyInfo(ctx context.Context, cafeId, userId int) (domain.Member, error) {
	reqUrl := fmt.Sprintf("%s/%d/info/%d", memberURL, cafeId, userId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("cli GetCafeMyInfo NewRequest err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("cli GetCafeMyInfo NewRequest err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return domain.Member{}, errors.New("internal server error")
	}

	var md domain.Member
	err = json.NewDecoder(resp.Body).Decode(&md)
	if err != nil {
		log.Println("GetCafeMyInfo json decode err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}
	return md, nil
}

func (r Requester) GetCafeIdsAndTotalByUserId(ctx context.Context, userId int, reqPage page.ReqPage) (domain.IdsTotalDomain, error) {
	reqUrl := fmt.Sprintf("%s/list/%d?page=%d&size=%d", memberURL, userId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId NewRequest err: ", err)
		return domain.IdsTotalDomain{}, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId DefaultClient.Do err: ", err)
		return domain.IdsTotalDomain{}, errors.New("internal server error")
	}
	defer resp.Body.Close()

	var iTD domain.IdsTotalDomain
	err = json.NewDecoder(resp.Body).Decode(&iTD)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId json.decode err: ", err)
		return domain.IdsTotalDomain{}, errors.New("internal server error")
	}
	return iTD, nil
}

func (r Requester) JoinCafe(ctx context.Context, d domain.Member) error {
	reqUrl := fmt.Sprintf("%s/%d/join/%d", memberURL, d.CafeId, d.UserId)
	jd := dto2.ToJoinMemberDto(d)
	data, err := json.Marshal(jd)
	if err != nil {
		log.Println("JoinCafe json.Marshal err: ", err)
		return errors.New("internal server error")
	}

	re, err := http.NewRequest("POST", reqUrl, bytes.NewReader(data))
	if err != nil {
		log.Println("JoinCafe NewRequest err: ", err)
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("JoinCafe DefaultClient.Do err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("JoinCafe readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) GetCafeMemberListCount(ctx context.Context, cafeId int, isBanned bool, reqPage page.ReqPage) (domain.MemberListCount, error) {
	reqUrl := fmt.Sprintf("%s/admin/%d?ban=%t&page=%d&size=%d", memberURL, cafeId, isBanned, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetCafeMemberListCount NewRequest err: ", err)
		return domain.MemberListCount{}, errors.New("internal server err")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetCafeMemberListCount defaultClient do err: ", err)
		return domain.MemberListCount{}, errors.New("internal server err")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetCafeMemberListCount readBody err: ", err)
			return domain.MemberListCount{}, errors.New("internal server err")
		}
		return domain.MemberListCount{}, errors.New(string(readBody))
	}
	var listCount domain.MemberListCount
	err = json.NewDecoder(resp.Body).Decode(&listCount)
	if err != nil {
		log.Println("GetCafeMemberListCount json decode err: ", err)
	}
	return listCount, nil
}

func (r Requester) PatchMember(ctx context.Context, d domain.Member) error {
	reqUrl := fmt.Sprintf("%s/%d/modify/%d", memberURL, d.CafeId, d.UserId)
	dto := dto2.ToPatchDto(d)
	data, err := json.Marshal(dto)
	if err != nil {
		log.Println("PatchMember json marshal err: ", err)
		return errors.New("internal server error")
	}

	re, err := http.NewRequest("PATCH", reqUrl, bytes.NewReader(data))
	if err != nil {
		log.Println("PatchMember NewRequest marshal err: ", err)
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("PatchMember DefaultClient marshal err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("PatchMember readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) GetMemberByCafeMemberId(ctx context.Context, cafeId int, memberId int) (domain.Member, error) {
	reqUrl := fmt.Sprintf("%s/admin/%d/%d", memberURL, cafeId, memberId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetMemberByCafeMemberId NewRequest err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetMemberByCafeMemberId DefaultClient.Do err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}
	defer resp.Body.Close()

	var dto dto2.MemberInfoDto
	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		log.Println("GetMemberByCafeMemberId json.NewDecoder err err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}
	mDomain := dto.ToDomain()
	return mDomain, nil
}
