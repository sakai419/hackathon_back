package models

type CreateAccountRequest struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type CreateAccountResponse struct {
	ID string `json:"id"`
}