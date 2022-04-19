package goconfluence

type GetAllUsersWithAnyPermissionType struct {
	MaxResults int64    `json:"maxResults"`
	StartAt    int64    `json:"startAt"`
	Total      int64    `json:"total"`
	Users      []string `json:"users"`
}
