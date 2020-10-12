package goconfluence

import (
	"net/http"
	"net/url"
)

// API is the main api data structure
type API struct {
	endPoint        *url.URL
	client          *http.Client
	username, token string
}

// Results array
type Results struct {
	ID      string  `json:"id,omitempty"`
	Type    string  `json:"type,omitempty"`
	Status  string  `json:"status,omitempty"`
	Content Content `json:"content"`
	Excerpt string  `json:"excerpt,omitempty"`
	Title   string  `json:"title,omitempty"`
	URL     string  `json:"url,omitempty"`
}

// Content specifies content properties
type Content struct {
	ID        string     `json:"id,omitempty"`
	Type      string     `json:"type,omitempty"`
	Status    string     `json:"status,omitempty"`
	Title     string     `json:"title,omitempty"`
	Ancestors []Ancestor `json:"ancestors"`
	Body      Body       `json:"body"`
	Version   Version    `json:"version"`
	Space     Space      `json:"space"`
}

// Ancestor defines ancestors to create sub pages
type Ancestor struct {
	ID string `json:"id"`
}

// Body holds the storage information
type Body struct {
	Storage Storage  `json:"storage"`
	View    *Storage `json:"view,omitempty"`
}

// Storage defines the storage information
type Storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

// Version defines the content version number
// the version number is used for updating content
type Version struct {
	Number    int    `json:"number"`
	MinorEdit bool   `json:"minorEdit"`
	Message   string `json:"message,omitempty"`
}

// Space holds the Space information of a Content Page
type Space struct {
	ID     int    `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
}

// ContentQuery defines the query parameters
// used for content related searching
// Query parameter values https://developer.atlassian.com/cloud/confluence/rest/#api-content-get
type ContentQuery struct {
	Expand     []string
	Limit      int    // page limit
	OrderBy    string // fieldpath asc/desc e.g: "history.createdDate desc"
	PostingDay string // required for blogpost type Format: yyyy-mm-dd
	SpaceKey   string
	Start      int    // page start
	Status     string // current, trashed, draft, any
	Title      string // required for page
	Trigger    string // viewed
	Type       string // page, blogpost
	Version    int    //version number when not lastest

}

// User defines user informations
type User struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	UserKey     string `json:"userKey"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

// Search results
type Search struct {
	Results []Results `json:"results"`
	Start   int       `json:"start,omitempty"`
	Limit   int       `json:"limit,omitempty"`
	Size    int       `json:"size,omitempty"`
}

// ContentSearch results
type ContentSearch struct {
	Results []Content `json:"results"`
	Start   int       `json:"start,omitempty"`
	Limit   int       `json:"limit,omitempty"`
	Size    int       `json:"size,omitempty"`
}

// SearchQuery defines query parameters used for searchng
// Query parameter values https://developer.atlassian.com/cloud/confluence/rest/#api-search-get
type SearchQuery struct {
	CQL                   string
	CQLContext            string
	IncludeArchivedSpaces bool
	Limit                 int
	Start                 int
}

// History contains object history information
type History struct {
	LastUpdated LastUpdated `json:"lastUpdated"`
	Latest      bool        `json:"latest"`
	CreatedBy   User        `json:"createdBy"`
	CreatedDate string      `json:"createdDate"`
}

// LastUpdated  contains information about the last update
type LastUpdated struct {
	By           User   `json:"by"`
	When         string `json:"when"`
	FriendlyWhen string `json:"friendlyWhen"`
	Message      string `json:"message"`
	Number       int    `json:"number"`
	MinorEdit    bool   `json:"minorEdit"`
	SyncRev      string `json:"syncRev"`
	ConfRev      string `json:"confRev"`
}

// Labels is the label containter type
type Labels struct {
	Labels []Label `json:"results"`
	Start  int     `json:"start,omitempty"`
	Limit  int     `json:"limit,omitempty"`
	Size   int     `json:"size,omitempty"`
}

// Label contains label information
type Label struct {
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
	ID     string `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
}

// Watchers is a list of Watcher
type Watchers struct {
	Watchers []Watcher `json:"results"`
	Start    int       `json:"start,omitempty"`
	Limit    int       `json:"limit,omitempty"`
	Size     int       `json:"size,omitempty"`
}

// Watcher contains information about watching users of a page
type Watcher struct {
	Type      string `json:"type"`
	Watcher   User   `json:"watcher"`
	ContentID int    `json:"contentId"`
}

// AllSpaces results
type AllSpaces struct {
	Results []Space `json:"results"`
	Start   int     `json:"start,omitempty"`
	Limit   int     `json:"limit,omitempty"`
	Size    int     `json:"size,omitempty"`
}

// AllSpacesQuery defines the query parameters
// Query parameter values https://developer.atlassian.com/cloud/confluence/rest/#api-space-get
type AllSpacesQuery struct {
	Expand           []string
	Favourite        bool   // Filter the results to the favourite spaces of the user specified by favouriteUserKey
	FavouriteUserKey string // The userKey of the user, whose favourite spaces are used to filter the results when using the favourite parameter. Leave blank for the current user
	Limit            int    // page limit
	SpaceKey         string
	Start            int    // page start
	Status           string // current, archived
	Type             string // global, personal
}
