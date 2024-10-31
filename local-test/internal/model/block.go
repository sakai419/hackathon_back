package model

import (
	"local-test/pkg/apperrors"
)

type BlockUserParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (p *BlockUserParams) Validate() error {
	if p.BlockedAccountID == p.BlockerAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "blocked account id and blocker account id must be different",
		}
	}

	return nil
}

type UnblockUserParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (p *UnblockUserParams) Validate() error {
	if p.BlockedAccountID == p.BlockerAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "blocked account id and blocker account id must be different",
		}
	}

	return nil
}

type GetBlockedAccountIDsParams struct {
	BlockerAccountID string
	Limit            int32
	Offset           int32
}

type GetBlockedInfosParams struct {
	BlockerAccountID string
	Limit            int32
	Offset           int32
}

func (p *GetBlockedInfosParams) Validate() error {
	if p.Limit < 1 {
		return &apperrors.ErrInvalidInput{
			Message: "limit must be greater than or equal to 1",
		}
	}
	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "offset must be greater than or equal to 0",
		}
	}

	return nil
}

type GetBlockerAccountIDsParams struct {
	ClientAccountID string
	IDs             []string
}

type IsBlockedParams struct {
	BlockerAccountID string
	BlockedAccountID string
}