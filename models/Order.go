package models

import "time"

type Order struct {
	OrderId    int       `json:"order_id"`
	ProductId  int       `json:"product_id"`
	StartedAt  time.Time `json:"started_at"`
	WillDoneAt time.Time `json:"will_done_at"`
	Status     bool      `json:"status"`
}
