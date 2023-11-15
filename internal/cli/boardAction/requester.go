package boardAction

import (
	"bytes"
	"cafe/internal/cli/boardAction/model"
	"cafe/internal/cli/boardAction/req"
	"cafe/internal/domain/boardAction"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8088/board-actions"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

const (
	InternalServerError = "internal server error"
)

func (r Requester) GetInfo(ctx context.Context, cafeId, boardTypeId int) (boardAction.BoardAction, error) {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, cafeId, boardTypeId)

	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetInfo NewRequest err: ", err)
		return boardAction.NewBuilder().Build(), errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		return boardAction.NewBuilder().Build(), errors.New(InternalServerError)
	}
	defer resp.Body.Close()
	var m model.BoardAction

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetInfo readBody err: ", err)
			return boardAction.NewBuilder().Build(), errors.New(InternalServerError)
		}
		return boardAction.NewBuilder().Build(), errors.New(string(readBody))
	}

	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Println("GetInfo json.NewDecoder err: ", err)
		return boardAction.NewBuilder().Build(), errors.New(InternalServerError)
	}
	return m.ToDomain(), nil
}

func (r Requester) Create(ctx context.Context, c req.Create) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, c.CafeId, c.BoardType)
	dto := c.ToCreateDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(dto)
	if err != nil {
		log.Println("Create NewEncoder err: ", err)
		return errors.New(InternalServerError)
	}

	re, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		log.Println("Create NewRequest err: ", err)
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
	reqUrl := fmt.Sprintf("%s/%d/%d/%d", baseUrl, p.CafeId, p.BoardTypeId, p.Id)
	dto := p.ToPatchDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(dto)
	if err != nil {
		log.Println("Patch json.NewEncoder err: ", err)
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

func (r Requester) Delete(ctx context.Context, cafeId int, boardTypeId int, id int) error {
	reqUrl := fmt.Sprintf("%s/%d/%d/%d", baseUrl, cafeId, boardTypeId, id)
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
