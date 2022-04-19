package goconfluence

type ContentVersionResult struct {
	Links struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Self    string `json:"self"`
	} `json:"_links"`
	Limit   int64 `json:"limit"`
	Results []struct {
		Expandable struct {
			Content string `json:"content"`
		} `json:"_expandable"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
		By struct {
			DisplayName    string `json:"displayName"`
			ProfilePicture struct {
				Height    int64  `json:"height"`
				IsDefault bool   `json:"isDefault"`
				Path      string `json:"path"`
				Width     int64  `json:"width"`
			} `json:"profilePicture"`
			Type string `json:"type"`
		} `json:"by"`
		Hidden    bool   `json:"hidden"`
		Message   string `json:"message"`
		MinorEdit bool   `json:"minorEdit"`
		Number    int64  `json:"number"`
		When      string `json:"when"`
	} `json:"results"`
	Size  int64 `json:"size"`
	Start int64 `json:"start"`
}
