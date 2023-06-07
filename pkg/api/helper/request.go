package helper

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type PaginationRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per-page"`
	Total   int `json:"total"`
}

type Ids struct {
	ID []int `json:"id"`
}
