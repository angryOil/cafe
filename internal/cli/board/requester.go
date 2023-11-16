package board

import (
	"bytes"
	"cafe/internal/cli/board/model"
	req2 "cafe/internal/cli/board/req"
	"cafe/internal/domain/board"
	page2 "cafe/internal/page"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8089/boards"

type Requester struct {
}

const (
	InternalServerError = "internal server error"
)

func (r Requester) GetList(ctx context.Context, l req2.GetList, reqPage page2.ReqPage) ([]board.Board, int, error) {
	reqUrl := fmt.Sprintf("%s/%d?board-type=%d&writer=%d&page=%d&size=%d", baseUrl, l.CafeId, l.BoardType, l.Writer, reqPage.Page, reqPage.Size)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []board.Board{}, 0, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DefaultClient.Do err: ", err)
		return []board.Board{}, 0, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetList readBody err: ", err)
			return []board.Board{}, 0, errors.New(InternalServerError)
		}
		return []board.Board{}, 0, errors.New(string(readBody))
	}

	var listTotalDto model.BoardListTotalDto
	err = json.NewDecoder(resp.Body).Decode(&listTotalDto)
	if err != nil {
		log.Println("GetList json.NewDecoder err: ", err)
		return []board.Board{}, 0, errors.New(InternalServerError)
	}
	return model.ToDomainList(listTotalDto.Content), listTotalDto.Total, nil
}

func (r Requester) Create(ctx context.Context, c req2.Create) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, c.CafeId, c.BoardType)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(c.ToCreateDto())
	if err != nil {
		log.Println("Create json.NewEncoder err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		log.Println("Create NewRequest err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Create DefaultClient.Do err: ", err)
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

func (r Requester) Patch(ctx context.Context, p req2.Patch) error {
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, p.Id)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p.ToPatchDto())
	if err != nil {
		log.Println("Patch NewEncoder err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("PATCH", reqUrl, &buf)
	if err != nil {
		log.Println("Patch NewRequest err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Patch DefaultClient.Do err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Patch readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}

func (r Requester) Delete(ctx context.Context, id int) error {
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, id)
	re, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		log.Println("Delete NewRequest err: ", err)
		return errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Delete DefaultClient.Do err: ", err)
		return errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Delete readBody err: ", err)
			return errors.New(InternalServerError)
		}
		return errors.New(string(readBody))
	}
	return nil
}

func NewRequester() Requester {
	return Requester{}
}
