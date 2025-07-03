package dtos

type CreateBlogDto struct {
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Content          string `json:"content"`
}
