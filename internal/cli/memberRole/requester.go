package memberRole

import (
	"bytes"
	"cafe/internal/cli/memberRole/cDto"
	"cafe/internal/domain"
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

func (r Requester) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.MemberRole, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?page=%d&size=%d", baseUrl, cafeId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetRolesByCafeId NewRequest err: ", err)
		return []domain.MemberRole{}, 0, errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetRolesByCafeId DefaultClient err: ", err)
		return []domain.MemberRole{}, 0, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetRolesByCafeId readBody err: ", err)
			return []domain.MemberRole{}, 0, errors.New("internal server error")
		}
		return []domain.MemberRole{}, 0, errors.New(string(readBody))
	}

	var listTotalDto cDto.ListTotalDto[cDto.MemberDetailRole]
	err = json.NewDecoder(resp.Body).Decode(&listTotalDto)
	if err != nil {
		log.Println("GetRolesByCafeId NewDecoder err: ", err)
		return []domain.MemberRole{}, 0, errors.New("internal server error")
	}

	return cDto.ToDomainList(listTotalDto.Contents), listTotalDto.Total, nil
}

func (r Requester) GetOneMemberRoles(ctx context.Context, cafeId int, memberId int) (domain.MemberRole, error) {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, memberId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetOneMemberRoles NewRequest err: ", err)
		return domain.MemberRole{}, errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetOneMemberRoles DefaultClient err: ", err)
		return domain.MemberRole{}, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetOneMemberRoles readBody err: ", err)
			return domain.MemberRole{}, errors.New("internal server error")
		}
		return domain.MemberRole{}, errors.New(string(readBody))
	}

	var d cDto.MemberRole
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Println("GetOneMemberRoles json.NewDecoder err: ", err)
		return domain.MemberRole{}, errors.New("internal server error")
	}
	return d.ToDomain(), nil
}

func (r Requester) PutRole(ctx context.Context, cafeId int, memberId int, d domain.MemberRole) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, memberId)
	cDto := cDto.ToPutMemberRoleCDto(d)
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
