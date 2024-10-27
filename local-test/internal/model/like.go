package model

type LikeTweetAndNotifyParams struct {
	LikingAccountID string
	OriginalTweetID int64
}

type CreateLikeAndNotifyParams struct {
	LikingAccountID string
	LikedAccountID  string
	OriginalTweetID int64
}

type UnlikeTweetParams struct {
	LikingAccountID string
	OriginalTweetID int64
}