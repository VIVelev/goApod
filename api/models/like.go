package models

// Like struct
type Like struct {
	AuthorID  int `json:"authorId"`
	ArticleID int `json:"articleId"`
}
