package req

type CreateBoardTypeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PatchBoardTypeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
