package book

type createRequest struct {
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Pages         int      `json:"pages"`
	ISBN          *string  `json:"isbn"`
	PublishedYear *int     `json:"published_year"`
	Description   *string  `json:"description"`
}

type deleteRequest struct {
	Id int64 `json:"id"`
}
