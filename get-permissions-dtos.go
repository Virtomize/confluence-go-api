package goconfluence

type AllSpaces2 struct {
	Links struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Self    string `json:"self"`
	} `json:"_links"`
	Limit   int64 `json:"limit"`
	Results []struct {
		Expandable struct {
			Description     string `json:"description"`
			Homepage        string `json:"homepage"`
			Icon            string `json:"icon"`
			Metadata        string `json:"metadata"`
			RetentionPolicy string `json:"retentionPolicy"`
		} `json:"_expandable"`
		Links struct {
			Self  string `json:"self"`
			Webui string `json:"webui"`
		} `json:"_links"`
		ID   int64  `json:"id"`
		Key  string `json:"key"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
	Size  int64 `json:"size"`
	Start int64 `json:"start"`
}
