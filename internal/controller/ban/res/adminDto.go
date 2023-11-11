package res

type BanAdminListDto struct {
	Id          int    `json:"id"`
	MemberId    int    `json:"member_id"`
	Description string `json:"description"`
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
