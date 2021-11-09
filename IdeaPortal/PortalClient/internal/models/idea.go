package models

type Idea struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	EstimatedTime int    `json:"estimatedTime"`
	CreatedDate   string `json:"createdData"`
}
