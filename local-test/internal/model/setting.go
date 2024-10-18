package model

type UpdateSettingsParams struct {
	AccountID        string  `json:"account_id"`
	IsPrivate        *bool   `json:"is_private"`
}