package boardAction

import (
	"cafe/internal/cli/boardAction/model"
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
