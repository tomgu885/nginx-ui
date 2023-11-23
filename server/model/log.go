package model

type Log struct {
	BaseModel
	Title   string `json:"title"`
	Content string `json:"content"`
}
