package req

type CreateBan struct {
	UserId      int
	MemberId    int
	CafeId      int
	Description string
}
