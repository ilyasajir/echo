package model

type Category struct {
	CategoryID   int    `json:"categoryid" form:"categoryid"`
	NamaCategory string `json:"namacategory" form:"namacategory"`
}
