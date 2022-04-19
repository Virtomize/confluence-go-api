package goconfluence

type GetPermissionsForSpaceType struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
