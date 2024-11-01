package model

type Paging struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
