package model

import "errors"

type BlockUserParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (p *BlockUserParams) Validate() error {
	if p.BlockedAccountID == p.BlockerAccountID {
		return errors.New("blocked account id and blocker account id must be different")
	}

	return nil
}

type UnblockUserParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (p *UnblockUserParams) Validate() error {
	if p.BlockedAccountID == p.BlockerAccountID {
		return errors.New("blocked account id and blocker account id must be different")
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
		return errors.New("limit must be greater than 0")
	}
	if p.Offset < 0 {
		return errors.New("offset must be greater than or equal to 0")
	}

	return nil
}