package goconfluence

type UsersType struct {
	MaxResults int64  `json:"maxResults"`
	StartAt    int64  `json:"startAt"`
	Status     string `json:"status"`
	Total      int64  `json:"total"`
	Users      []struct {
		Business []struct {
			Department string `json:"department"`
			Location   string `json:"location"`
			Position   string `json:"position"`
		} `json:"business"`
		CreatedDate                   int64       `json:"createdDate"`
		CreatedDateString             string      `json:"createdDateString"`
		Email                         string      `json:"email"`
		FullName                      string      `json:"fullName"`
		HasAccessToUseConfluence      bool        `json:"hasAccessToUseConfluence"`
		Key                           string      `json:"key"`
		LastFailedLoginDate           interface{} `json:"lastFailedLoginDate"`
		LastFailedLoginDateString     interface{} `json:"lastFailedLoginDateString"`
		LastSuccessfulLoginDate       int64       `json:"lastSuccessfulLoginDate"`
		LastSuccessfulLoginDateString string      `json:"lastSuccessfulLoginDateString"`
		Name                          string      `json:"name"`
		Personal                      []struct {
			Im      string `json:"im"`
			Phone   string `json:"phone"`
			Website string `json:"website"`
		} `json:"personal"`
		UpdatedDate       int64  `json:"updatedDate"`
		UpdatedDateString string `json:"updatedDateString"`
	} `json:"users"`
}
