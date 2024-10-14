package model

import "local-test/pkg/utils"

type CreateAccountParams struct {
	ID       string
	UserID   string
	UserName string
}

func (p *CreateAccountParams) Validate() error {
	if len(p.ID) != 28 {
		return &utils.ErrInvalidInput{Message: "invalid firebase uid"}
	}
	if len(p.UserID) > 30 {
		return &utils.ErrInvalidInput{Message: "user id is too long"}
	}
	if len(p.UserName) > 30 {
		return &utils.ErrInvalidInput{Message: "user name is too long"}
	}
	return nil
}

type GetUserAndProfileInfoByAccountIDsParams struct {
	Limit  int32
	Offset int32
	IDs    []string
}

type UserAndProfileInfo struct {
	UserID          string
	UserName        string
	Bio	            string
	ProfileImageURL string
}