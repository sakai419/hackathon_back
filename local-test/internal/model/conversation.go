package model

import "time"

type GetConversationIDParams struct {
	Account1ID    string
	Account2ID  string
}

type GetConversationsParams struct {
	ClientAccountID string
	Limit     int32
	Offset    int32
}

type GetConversationListParams struct {
	AccountID string
	Limit     int32
	Offset    int32
}

type Conversation struct {
	ID              int64
	OpponentID      string
	LastMessageTime time.Time
	Content         string
	SenderAccountID string
	IsRead          bool
}

type ConversationResponse struct {
	ID              int64
	OpponentInfo    UserInfoWithoutBio
	LastMessageTime time.Time
	Content         string
	SenderID        string
	IsRead	        bool
}