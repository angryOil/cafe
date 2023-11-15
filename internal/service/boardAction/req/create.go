package req

type Create struct {
	CafeId      int
	BoardType   int
	ReadRoles   string
	CreateRoles string
	UpdateRoles string
	UpdateAble  bool
	DeleteRoles string
}
