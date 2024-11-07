package model

type SidebarInfo struct {
	UserInfo UserInfoWithoutBio
	UnreadConversationCount int64
	UnreadNotificationCount int64
}