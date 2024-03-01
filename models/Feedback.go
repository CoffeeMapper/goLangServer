package models

type Feedback struct {
	FeedbackId int    `json:"feedback_id"`
	UserId     int    `json:"user_id"`
	Feedback   string `json:"feedback"`
	Rate       int    `json:"rate"`
}
