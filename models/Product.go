package models

type Product struct {
	ProductId   int    `json:"product_id"`
	CategoryId  int    `json:"category_id"`
	Title       string `json:"title"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
