package model

type RetweetAndNotifyParams struct {
	RetweetingAccountID string
	OriginalTweetID	    int64
}

type CreateRetweetAndNotifyParams struct {
	RetweetingAccountID string
	RetweetedAccountID  string
	OriginalTweetID     int64
}

type UnretweetParams struct {
	RetweetingAccountID string
	OriginalTweetID     int64
}