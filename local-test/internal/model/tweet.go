package model

import (
	"local-test/pkg/apperrors"
	"time"
)

const (
	MediaTypeJPG = "jpg"
)

type Media struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (m *Media) Validate() error {
	switch m.Type {
	case MediaTypeJPG:
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

type PostTweetParams struct {
	AccountID       string
	Content         *string
	Code            *string
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

type CreateTweetParams struct {
	AccountID   string
	Content     *string
	Code        *string
	Media       *Media
	HashtagIDs  []int64
}

type GetTweetLabelsParams struct {
	Content *string
	Code    *string
	Media   *Media
}

type GetTweetInfosByAccountIDParams struct {
	ClientAccountID string
	TargetAccountID string
	Limit	  	    int32
	Offset	 	    int32
}

type TweetInfoInternal struct {
	TweetID       int64
	AccountID     string
	Content       *string
	Code          *string
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
	Code          *string
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