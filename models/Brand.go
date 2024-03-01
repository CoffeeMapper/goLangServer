package models

type Brand struct {
	BrandId int    `json:"brand_id"`
	Name    string `json:"name"`
	OwnerId int    `json:"owner_id"`
	Phone   string `json:"phone"`
}
