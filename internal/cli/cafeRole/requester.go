package cafeRole

import (
	"bytes"
	"cafe/internal/cli/cafeRole/model"
	"cafe/internal/cli/cafeRole/req"
	"cafe/internal/domain/cafeRole"
	page2 "cafe/internal/page"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8086/roles"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

func (r Requester) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]cafeRole.CafeRole, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?page=%d&size=%d", baseUrl, cafeId, reqPage.Page, reqPage.Size)

	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []cafeRole.CafeRole{}, 0, errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DefaultClient err: ", err)
		return []cafeRole.CafeRole{}, 0, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetList readBody err: ", err)
			return []cafeRole.CafeRole{}, 0, errors.New("internal server error")
		}
		return []cafeRole.CafeRole{}, 0, errors.New(string(readBody))
	}

	var listTotalDto model.CafeRoleListTotal
	err = json.NewDecoder(resp.Body).Decode(&listTotalDto)
	if err != nil {
		log.Println("GetList json.NewDecoder err: ", err)
		return []cafeRole.CafeRole{}, 0, errors.New("internal server error")
	}
	domains := make([]cafeRole.CafeRole, len(listTotalDto.Roles))
	for i, l := range listTotalDto.Roles {
		domains[i] = l.ToDomain()
	}
	return domains, listTotalDto.Total, nil
}

func (r Requester) Create(ctx context.Context, c req.Create) error {
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, c.CafeId)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(c.ToCreateDto())
	if err != nil {
		log.Println("Create json.NewEncoder err: ", err)
		return errors.New("internal server error")
	}
	re, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		log.Println("Create NewRequest err: ", err)
		return errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Create http.DefaultClient err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Create readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) Patch(ctx context.Context, p req.Patch) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, p.CafeId, p.Id)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p.ToPatchDto())
	if err != nil {
		log.Println("Patch json.NewEncoder err: ", err)
		return errors.New("internal server error")
	}
	re, err := http.NewRequest("PATCH", reqUrl, &buf)
	if err != nil {
		log.Println("Patch http.NewRequest err: ", err)
		return errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Patch http.DefaultClient err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Patch readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) Delete(ctx context.Context, cafeId int, roleId int) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, roleId)
	re, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		log.Println("Delete NewRequest err: ", err)
		return errors.New("internal server error")
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Delete DefaultClient err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Delete readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}
