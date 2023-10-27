package member

import (
	"cafe/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// memberURL
var memberURL = "http://localhost:8084/members"

type Requester struct {
}

func NewRequester() Requester {
	return Requester{}
}

// 해당 카페의 member(자기정보) 정보조회

func (r Requester) GetCafeMyInfo(ctx context.Context, cafeId, userId int) (domain.Member, error) {
	reqUrl := fmt.Sprintf("%s/%d/info/%d", memberURL, cafeId, userId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("cli GetCafeMyInfo NewRequest err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("cli GetCafeMyInfo NewRequest err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return domain.Member{}, errors.New("internal server error")
	}

	var md domain.Member
	err = json.NewDecoder(resp.Body).Decode(&md)
	if err != nil {
		log.Println("GetCafeMyInfo json decode err: ", err)
		return domain.Member{}, errors.New("internal server error")
	}
	return md, nil
}

func (r Requester) GetCafeIdsAndTotalByUserId(ctx context.Context, userId int) (domain.IdsTotalDomain, error) {
	reqUrl := fmt.Sprintf("%s/list/%d", memberURL, userId)
	re, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId NewRequest err: ", err)
		return domain.IdsTotalDomain{}, errors.New("internal server error")
	}

	resp, err := http.DefaultClient.Do(re)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId DefaultClient.Do err: ", err)
		return domain.IdsTotalDomain{}, errors.New("internal server error")
	}
	defer resp.Body.Close()

	var iTD domain.IdsTotalDomain
	err = json.NewDecoder(resp.Body).Decode(&iTD)
	if err != nil {
		log.Println("GetCafeIdsAndTotalByUserId json.decode err: ", err)
		return domain.IdsTotalDomain{}, errors.New("internal server error")
	}
	return iTD, nil
}
