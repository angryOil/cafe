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

var boardTypeUrl = "http://localhost:8085/board-types"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

func (r Requester) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]boardType.BoardType, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?page=%d&size=%d", boardTypeUrl, cafeId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []boardType.BoardType{}, 0, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DefaultClient err: ", err)
		return []boardType.BoardType{}, 0, errors.New("internal server error")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetList readBody err: ", err)
			return []boardType.BoardType{}, 0, errors.New("internal server error")
		}
		return []boardType.BoardType{}, 0, errors.New(string(readBody))
	}

	var listCount model.BoardTypeListCount
	err = json.NewDecoder(resp.Body).Decode(&listCount)
	if err != nil {
		log.Println("GetList json.NewDecoder err: ", err)
		return []boardType.BoardType{}, 0, errors.New("internal server error")
	}

	domains := make([]boardType.BoardType, len(listCount.BoardTypes))
	for i, b := range listCount.BoardTypes {
		domains[i] = boardType.NewBuilder().
			Id(b.Id).
			Name(b.Name).
			Description(b.Description).
			Build()
	}
	return domains, listCount.Total, nil
}

func (r Requester) Create(ctx context.Context, c req.Create) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", boardTypeUrl, c.CafeId, c.OwnerId)
	cDto := c.ToCreateDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(cDto)
	if err != nil {
		log.Println("Create json.NewEncode err: ", err)
		return errors.New("internal server error")
	}

	re, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		log.Println("Create NewRequest err")
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Create DefaultClient err: ", err)
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
	reqUrl := fmt.Sprintf("%s/%d/%d", boardTypeUrl, p.CafeId, p.Id)

	pDto := p.ToPatchDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(pDto)
	if err != nil {
		log.Println("patch json.NewEncode err: ", err)
		return errors.New("internal server error")
	}

	re, err := http.NewRequest("PATCH", reqUrl, &buf)
	if err != nil {
		log.Println("patch NewRequest err: ", err)
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("patch DefaultClient err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("patch readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) Delete(ctx context.Context, cafeId int, typeId int) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", boardTypeUrl, cafeId, typeId)

	re, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		log.Println("delete NewRequest err: ", err)
		return errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("delete DefaultClient err: ", err)
		return errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("delete readBody err: ", err)
			return errors.New("internal server error")
		}
		return errors.New(string(readBody))
	}
	return nil
}
