package model

type LeftSidebarInfo struct {
	UnreadConversationCount int64
	UnreadNotificationCount int64
}

type RightSidebarInfo struct {
	RecentLabels []*LabelCount
	FollowSuggestions []*UserInfoWithoutBio
}