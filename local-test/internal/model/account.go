package model

// CreateAccount
type CreateAccountServiceParams struct {
	ID       string
	UserID   string
	UserName string
}

func (p *CreateAccountServiceParams) ToParams() *CreateAccountRepositoryParams {
	return &CreateAccountRepositoryParams{
		ID:       p.ID,
		UserID:   p.UserID,
		UserName: p.UserName,
	}
}

type CreateAccountRepositoryParams struct {
	ID       string
	UserID   string
	UserName string
}

// DeleteMyAccount
type DeleteMyAccountServiceParams struct {
	ID string
}

func (p *DeleteMyAccountServiceParams) ToParams() *DeleteMyAccountRepositoryParams {
	return &DeleteMyAccountRepositoryParams{
		ID: p.ID,
	}
}

type DeleteMyAccountRepositoryParams struct {
	ID string
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