package models

type Notification struct {
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
