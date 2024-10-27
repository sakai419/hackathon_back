package model

type CreateRetweetAndNotifyParams struct {
	RetweetingAccountID string
	RetweetedAccountID  string
	OriginalTweetID     int64
}

