package boardType

import (
	"bytes"
	"cafe/internal/cli/boardType/dto"
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

var boardTypeUrl = "http://localhost:8085/board-types"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

func (r Requester) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.BoardType, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?page=%d&size=%d", boardTypeUrl, cafeId, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []domain.BoardType{}, 0, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DefaultClient err: ", err)
		return []domain.BoardType{}, 0, errors.New("internal server error")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetList readBody err: ", err)
			return []domain.BoardType{}, 0, errors.New("internal server error")
		}
		return []domain.BoardType{}, 0, errors.New(string(readBody))
	}

	var listCount dto.BoardTypeCDtoListCount
	err = json.NewDecoder(resp.Body).Decode(&listCount)
	if err != nil {
		log.Println("GetList json.NewDecoder err: ", err)
		return []domain.BoardType{}, 0, errors.New("internal server error")
	}
	return dto.ToDomainList(listCount.Boards), listCount.Total, nil
}

func (r Requester) Create(ctx context.Context, typeDomain domain.BoardType) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", boardTypeUrl, typeDomain.CafeId, typeDomain.CreateBy)
	cDto := dto.ToCreateBoardTypeCDto(typeDomain)
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
