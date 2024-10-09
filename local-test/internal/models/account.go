package models

import "errors"

// CreateAccount
type CreateAccountRequest struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func (r *CreateAccountRequest) ToParams() *CreateAccountParams {
	return &CreateAccountParams{
		ID:       r.ID,
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
	if r.ID == "" {
		return errors.New("ID is required")
	}
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