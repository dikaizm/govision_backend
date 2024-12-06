package response

type GetArticle struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body,omitempty"`
	Image     string `json:"image"`
	ReadCount int    `json:"read_count"`
	Author    string `json:"author,omitempty"`
	CreatedAt string `json:"created_at"`
}
