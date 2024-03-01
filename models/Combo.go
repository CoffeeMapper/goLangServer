package models

type Combo struct {
	ComboId   int `json:"combo_id"`
	Title     int `json:"title"`
	Price     int `json:"price"`
	ProductId int `json:"product_id"`
}
