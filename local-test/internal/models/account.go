package models

import "errors"

// CreateAccount
type CreateAccountRequest struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func (r *CreateAccountRequest) ToParams() *CreateAccountParams {
	return &CreateAccountParams{
		ID:       "",
		UserID:   r.UserID,
		UserName: r.UserName,
	}
}

type CreateAccountParams struct {
	ID       string
	UserID   string
	UserName string
}

func (r *CreateAccountRequest) Validate() error {
	if r.UserID == "" {
		return errors.New("UserID is required")
	}
	if r.UserName == "" {
		return errors.New("UserName is required")
	}
	return nil
}

type CreateAccountResponse struct {
	ID string `json:"id"`
}