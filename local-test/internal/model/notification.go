package model

type NotificationsType string

const (
	NotificationsTypeFollow        NotificationsType = "follow"
	NotificationsTypeLike          NotificationsType = "like"
	NotificationsTypeRetweet       NotificationsType = "retweet"
	NotificationsTypeReply         NotificationsType = "reply"
	NotificationsTypeMessage       NotificationsType = "message"
	NotificationsTypeQuote         NotificationsType = "quote"
	NotificationsTypeFollowRequest NotificationsType = "follow_request"
	NotificationsTypeReport        NotificationsType = "report"
)
