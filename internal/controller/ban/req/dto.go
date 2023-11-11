package req

type CreateBanDto struct {
	MemberId    int    `json:"member_id"`
	Description string `json:"description"`
}
