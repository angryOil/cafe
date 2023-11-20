package boardType

import (
	"bytes"
	"cafe/internal/cli/boardType/model"
	"cafe/internal/cli/boardType/req"
	"cafe/internal/domain/boardType"
	page2 "cafe/internal/page"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8085/board-types"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

const (
	InternalServerError = "internal server error"
)

func (r Requester) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]boardType.BoardType, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?page=%d&size=%d", baseUrl, cafeId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []boardType.BoardType{}, 0, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DefaultClient err: ", err)
		return []boardType.BoardType{}, 0, errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetList readBody err: ", err)
			return []boardType.BoardType{}, 0, errors.New(InternalServerError)
		}
		return []boardType.BoardType{}, 0, errors.New(string(readBody))
	}

	var listCount model.BoardTypeListCount
	err = json.NewDecoder(resp.Body).Decode(&listCount)
	if err != nil {
		log.Println("GetList json.NewDecoder err: ", err)
		return []boardType.BoardType{}, 0, errors.New(InternalServerError)
	}

	domains := make([]boardType.BoardType, len(listCount.BoardTypes))
	for i, b := range listCount.BoardTypes {
		domains[i] = b.ToDomain()
	}
	return domains, listCount.Total, nil
}

func (r Requester) GetDetail(ctx context.Context, cafeId int, id int) (boardType.BoardType, error) {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, id)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetDetail NewRequest err: ", err)
		return boardType.NewBuilder().Build(), errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetDetail DefaultClient.Do err: ", err)
		return boardType.NewBuilder().Build(), errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	var m model.BoardType
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Println("GetDetail json.NewDecoder err: ", err)
		return boardType.NewBuilder().Build(), errors.New(InternalServerError)
	}
	return m.ToDomain(), nil
}

func (r Requester) Create(ctx context.Context, c req.Create) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, c.CafeId, c.OwnerId)
	cDto := c.ToCreateDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(cDto)
	if err != nil {
		log.Println("Create json.NewEncode err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		log.Println("Create NewRequest err")
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Create DefaultClient err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Create readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) Patch(ctx context.Context, p req.Patch) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, p.CafeId, p.Id)

	pDto := p.ToPatchDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(pDto)
	if err != nil {
		log.Println("patch json.NewEncode err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("PATCH", reqUrl, &buf)
	if err != nil {
		log.Println("patch NewRequest err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("patch DefaultClient err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("patch readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) Delete(ctx context.Context, cafeId int, typeId int) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, typeId)

	re, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		log.Println("delete NewRequest err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("delete DefaultClient err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("delete readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}
