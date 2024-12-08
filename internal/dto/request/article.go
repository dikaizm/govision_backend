package request

type CreateArticle struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	Image    string `json:"image" validate:"required"`
	AuthorID string
}

type FilterGetArticle struct {
	SizeParam string `schema:"size"`
	Size      int
}
