package book

type listResponse struct {
	CurrentPage int    `json:"page"`
	TotalPages  int    `json:"total_pages"`
	TotalBooks  int64  `json:"total_books"`
	Data        []Book `json:"data"`
}
