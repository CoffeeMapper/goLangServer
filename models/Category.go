package models

type Category struct {
	CategoryId  int    `json:"category_id"`
	Title       string `json:"title"`
	description string `json:"description"`
}
