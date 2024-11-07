package model

type SidebarInfo struct {
	UserInfo UserInfoWithoutBio
	UnreadConversationCount int64
	UnReadNotificationCount int64
}