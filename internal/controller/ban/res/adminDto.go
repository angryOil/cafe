package res

import "cafe/internal/domain"

type BanAdminListDto struct {
	Id          int    `json:"id"`
	MemberId    int    `json:"member_id"`
	Description string `json:"description"`
}

func ToAdminDtoList(domains []domain.Ban) []BanAdminListDto {
	result := make([]BanAdminListDto, len(domains))
	for i, d := range domains {
		result[i] = BanAdminListDto{
			Id:          d.Id,
			MemberId:    d.MemberId,
			Description: d.Description,
		}
	}
	return result
}

type BanAdminDetailDto struct {
	Id             int    `json:"id"`
	MemberNickname string `json:"member_nickname"`
	MemberId       int    `json:"member_id"`
	Description    string `json:"description"`
}

func (d BanAdminListDto) ToDetailDto(nickname string) BanAdminDetailDto {
	return BanAdminDetailDto{
		Id:             d.Id,
		MemberNickname: nickname,
		MemberId:       d.MemberId,
		Description:    d.Description,
	}
}
