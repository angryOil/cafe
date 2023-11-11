package res

type BanListDto struct {
	Id          int    `json:"id"`
	CafeId      int    `json:"cafe_id"`
	Description string `json:"description"`
}

type BanDetailDto struct {
	Id          int    `json:"id"`
	CafeId      int    `json:"cafe_id"`
	CafeName    string `json:"cafe_name"`
	Description string `json:"description"`
}

func (d BanListDto) ToDetailDto(cafeName string) BanDetailDto {
	return BanDetailDto{
		Id:          d.Id,
		CafeId:      d.CafeId,
		CafeName:    cafeName,
		Description: d.Description,
	}
}
