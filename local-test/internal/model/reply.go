package model

import "local-test/pkg/apperrors"

type PostReplyAndNotifyParams struct {
	ReplyingAccountID string
	OriginalTweetID   int64
	Content           *string
	Code 			  *Code
	Media             *Media
}

func (p *PostReplyAndNotifyParams) Validate() error {
	if p.Content == nil && p.Code == nil && p.Media == nil {
		return &apperrors.ErrInvalidInput{
			Message: "content, code, and media are all missing",
		}
	}

	if p.Media != nil {
		if err := p.Media.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type CreateReplyAndNotifyParams struct {
	ReplyingAccountID string
	RepliedAccountID  string
	OriginalTweetID   int64
	Content           *string
	Code 			  *Code
	Media             *Media
	HashtagIDs		  []int64
}

type GetReplyIDsParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset		    int32
}

type GetRepliedTweetInfosParams struct {
	ClientAccountID  string
	ReplyingTweetIDs []int64
}

type RepliedTweetInfoInternal struct {
	OriginalTweet     TweetInfoInternal
	ParentReplyTweet  *TweetInfoInternal
	ReplyingTweetID   int64
	OmittedReplyExist *bool
}