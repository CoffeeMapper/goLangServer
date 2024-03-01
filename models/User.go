package models

type User struct {
	UserId   int    `json:"user_id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	SubId    bool   `json:"has_premium"`
}
