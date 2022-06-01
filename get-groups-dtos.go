package goconfluence

type GroupsType struct {
	Groups     []string `json:"groups"`
	MaxResults int64    `json:"maxResults"`
	StartAt    int64    `json:"startAt"`
	Status     string   `json:"status"`
	Total      int64    `json:"total"`
}
