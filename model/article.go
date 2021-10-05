package model

type Article struct {
	ArticleID int    `json:"articleid" form:"articleid"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Createdat int    `json:"createdat" form:"createdat"`
}
