package memberRole

import (
	"bytes"
	"cafe/internal/cli/memberRole/model"
	"cafe/internal/cli/memberRole/req"
	"cafe/internal/domain/memberRole"
	page2 "cafe/internal/page"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8087/member-roles"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

func (r Requester) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]memberRole.MemberRole, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?page=%d&size=%d", baseUrl, cafeId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetRolesByCafeId NewRequest err: ", err)
		return []memberRole.MemberRole{}, 0, errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetRolesByCafeId DefaultClient err: ", err)
		return []memberRole.MemberRole{}, 0, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetRolesByCafeId readBody err: ", err)
			return []memberRole.MemberRole{}, 0, errors.New("internal server error")
		}
		return []memberRole.MemberRole{}, 0, errors.New(string(readBody))
	}

	var listTotalDto model.ListTotalDto
	err = json.NewDecoder(resp.Body).Decode(&listTotalDto)
	if err != nil {
		log.Println("GetRolesByCafeId NewDecoder err: ", err)
		return []memberRole.MemberRole{}, 0, errors.New("internal server error")
	}

	return model.ToDomainList(listTotalDto.Contents), listTotalDto.Total, nil
}

func (r Requester) GetOneMemberRoles(ctx context.Context, cafeId int, memberId int) (memberRole.MemberRole, error) {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, memberId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetOneMemberRoles NewRequest err: ", err)
		return memberRole.NewBuilder().Build(), errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetOneMemberRoles DefaultClient err: ", err)
		return memberRole.NewBuilder().Build(), errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetOneMemberRoles readBody err: ", err)
			return memberRole.NewBuilder().Build(), errors.New("internal server error")
		}
		return memberRole.NewBuilder().Build(), errors.New(string(readBody))
	}

	var m model.MemberRole
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Println("GetOneMemberRoles json.NewDecoder err: ", err)
		return memberRole.NewBuilder().Build(), errors.New("internal server error")
	}
	return m.ToDomain(), nil
}

func (r Requester) PutRole(ctx context.Context, p req.PutRole) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, p.CafeId, p.MemberId)
	cDto := p.ToDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(cDto)
	if err != nil {
		log.Println("PutRole json.NewEncoder err: ", err)
		return errors.New("internal server error")
	}

	re, err := http.NewRequest("PUT", reqUrl, &buf)
	if err != nil {
		log.Println("PutRole http.NewRequest err: ", err)
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("PutRole http.DefaultClient err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("PutRole readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}

const (
	InternalServerErr = "internal server error"
)

func (r Requester) Delete(ctx context.Context, cafeId int, memberId int, mRoleId int) error {
	reqUrl := fmt.Sprintf("%s/%d/%d/%d", baseUrl, cafeId, memberId, mRoleId)
	re, err := http.NewRequest("DELETE", reqUrl, nil)

	if err != nil {
		log.Println("Delete NewRequest err: ", err)
		return errors.New(InternalServerErr)
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Delete DefaultClient err: ", err)
		return errors.New(InternalServerErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}
