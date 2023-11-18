package reply

import (
	"cafe/internal/cli/reply/model"
	"cafe/internal/domain/reply"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl = "http://localhost:8090/replies"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

const (
	InternalServerError = "internal server error"
)

func (r Requester) GetList(ctx context.Context, boardId int) ([]reply.Reply, int, error) {
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, boardId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetList NewRequest err: ", err)
		return []reply.Reply{}, 0, errors.New(InternalServerError)
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetList DEfaultClient.Do err: ", err)
		return []reply.Reply{}, 0, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetList readBody err: ", err)
			return []reply.Reply{}, 0, errors.New(InternalServerError)
		}
		return []reply.Reply{}, 0, errors.New(string(readBody))
	}

	var listTotalDto model.ListTotalDto
	err = json.NewDecoder(resp.Body).Decode(&listTotalDto)
	if err != nil {
		log.Println("GetList json.NewDecoder err: ", err)
		return []reply.Reply{}, 0, errors.New(InternalServerError)
	}
	return model.ToDomainList(listTotalDto.Content), listTotalDto.Total, nil
}
