package reply

import (
	"bytes"
	"cafe/internal/cli/reply/model"
	"cafe/internal/cli/reply/req"
	"cafe/internal/domain/reply"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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

func (r Requester) GetCount(ctx context.Context, arr []int) ([]model.GetCount, error) {
	str := arrayToString(arr, ",")
	reqUrl := fmt.Sprintf("%s/cnt?board-ids=%s", baseUrl, str)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetCount NewRequest err: ", err)
		return []model.GetCount{}, errors.New(InternalServerError)
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetCount http.DefaultClient.Do err: ", err)
		return []model.GetCount{}, errors.New(InternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("GetCount readBody err: ", err)
			return []model.GetCount{}, errors.New(InternalServerError)
		}
		return []model.GetCount{}, errors.New(string(readBody))
	}

	var countListDto model.CountListDto
	err = json.NewDecoder(resp.Body).Decode(&countListDto)
	if err != nil {
		log.Println("GetCount json.NewDecoder err: ", err)
		return []model.GetCount{}, errors.New(InternalServerError)
	}
	return countListDto.Content, nil
}

func (r Requester) Create(ctx context.Context, c req.Create) error {
	reqUrl := fmt.Sprintf("%s/%d/%d", baseUrl, c.CafeId, c.BoardId)
	dto := c.ToCreateDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(dto)
	if err != nil {
		log.Println("Crete json.NewEncoder err: ", err)
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

func (r Requester) Delete(ctx context.Context, replyId int) error {
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, replyId)
	re, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		log.Println("Delete NewRequest err: ", err)
		return errors.New(InternalServerError)
	}
	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("Delete DefaultClient err: ", err)
		return errors.New(InternalServerError)
	}
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

func (r Requester) Patch(ctx context.Context, p req.Patch) error {
	reqUrl := fmt.Sprintf("%s/%d", baseUrl, p.Id)
	dto := p.ToPatchDto()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(dto)
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
		log.Println("Patch DefaultClient err: ", err)
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

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
