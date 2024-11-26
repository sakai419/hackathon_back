package model

import (
	"local-test/pkg/apperrors"
	"time"
)

const (
	MediaTypeImage = "image"
	MediaTypeVideo = "video"
)

type Media struct {
	Type string
	URL  string
}

func (m *Media) Validate() error {
	switch m.Type {
	case MediaTypeImage:
	case MediaTypeVideo:
		if m.URL == "" {
			return &apperrors.ErrInvalidInput{
				Message: "media url is missing",
			}
		}
	default:
		return &apperrors.ErrInvalidInput{
			Message: "media type is invalid",
		}
	}

	return nil
}

type Code struct {
	Language string
	Content  string
}

type PostTweetParams struct {
	AccountID       string
	Content         *string
	Code            *Code
	Media           *Media
}

func (p *PostTweetParams) Validate() error {
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

type SetTweetAsPinnedParams struct {
	TweetID         int64
	ClientAccountID string
}

type UnsetTweetAsPinnedParams struct {
	TweetID         int64
	ClientAccountID string
}

type CreateTweetParams struct {
	AccountID   string
	Content     *string
	Code        *Code
	Media       *Media
	HashtagIDs  []int64
}

type GetTweetLabelsParams struct {
	Content *string
	Code    *Code
	Media   *Media
}

type GetTweetInfosByAccountIDParams struct {
	ClientAccountID string
	TargetAccountID string
	Limit	  	    int32
	Offset	 	    int32
}

type GetTweetInfosByIDsParams struct {
	ClientAccountID string
	TweetIDs        []int64
}

type TweetInfoInternal struct {
	TweetID       int64
	AccountID     string
	Content       *string
	Code          *Code
	Media         *Media
	LikesCount    int32
	RetweetsCount int32
	RepliesCount  int32
	IsQuote	      bool
	IsReply	      bool
	IsPinned	  bool
	HasLiked	  bool
	HasRetweeted  bool
	CreatedAt     time.Time
}

type TweetInfo struct {
	TweetID       int64
	UserInfo	  UserInfoWithoutBio
	Content       *string
	Code          *Code
	Media         *Media
	LikesCount    int32
	RetweetsCount int32
	RepliesCount  int32
	IsQuote	      bool
	IsReply	      bool
	IsPinned	  bool
	HasLiked	  bool
	HasRetweeted  bool
	CreatedAt     time.Time
}

type GetReplyTweetInfosParams struct {
	ClientAccountID string
	ParentTweetID   int64
	Limit           int32
	Offset          int32
}

func (p *GetReplyTweetInfosParams) Validate() error {
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

type GetTimelineTweetInfosParams struct {
	ClientAccountID string
	Limit           int32
	Offset          int32
}

func (p *GetTimelineTweetInfosParams) Validate() error {
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

type GetTimelineTweetInfosResponse struct {
	Tweet             TweetInfo
	OriginalTweet     *TweetInfo
	ParentReply       *TweetInfo
	OmittedReplyExist *bool
}

type GetRecentTweetMetadatasParams struct {
	Limit  int32
	Offset int32
	ClientAccountID string
}

type DeleteTweetParams struct {
	TweetID         int64
	ClientAccountID string
}

type TweetMetadata struct {
	TweetID       int64
	AccountID     string
	LikesCount    int32
	RetweetsCount int32
	RepliesCount  int32
	Label1		  *Label
	Label2		  *Label
	Label3		  *Label
}