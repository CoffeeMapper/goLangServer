package models

type CoffeePoint struct {
	CoffeePointId int    `json:"coffee_point_id"`
	Logo          string `json:"logo"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	BrandId       int    `json:"brand_id"`
	WorkTime      string `json:"work_time"`
	OrderId       int    `json:"order_id"`
	AvgRating     int    `json:"avg_rating"`
	FeedbackId    int    `json:"feedback_id"`
}
