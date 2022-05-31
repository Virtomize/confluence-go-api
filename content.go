package goconfluence

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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
	Type      string     `json:"type"`
	Status    string     `json:"status,omitempty"`
	Title     string     `json:"title"`
	Ancestors []Ancestor `json:"ancestors,omitempty"`
	Body      Body       `json:"body"`
	Version   *Version   `json:"version,omitempty"`
	Space     Space      `json:"space"`
	History   *History   `json:"history,omitempty"`
	Links     *Links     `json:"_links,omitempty"`
}

// Links contains link information
type Links struct {
	Base   string `json:"base"`
	TinyUI string `json:"tinyui"`
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

// BodyExportView holds the export_view information
type BodyExportView struct {
	ExportView *Storage `json:"export_view"`
	View       *Storage `json:"view,omitempty"`
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
	By        *User  `json:"by,omitempty"`
}

// Space holds the Space information of a Content Page
type Space struct {
	ID     int    `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
}

// getContentIDEndpoint creates the correct api endpoint by given id
func (a *API) getContentIDEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/api/content/" + id)
}

// getContentEndpoint creates the correct api endpoint
func (a *API) getContentEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/api/content/")
}

// getContentChildEndpoint creates the correct api endpoint by given id and type
func (a *API) getContentChildEndpoint(id string, t string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/api/content/" + id + "/child/" + t)
}

// getContentGenericEndpoint creates the correct api endpoint by given id and type
func (a *API) getContentGenericEndpoint(id string, t string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/api/content/" + id + "/" + t)
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

// GetContentByID querys content by id
func (a *API) GetContentByID(id string, query ContentQuery) (*Content, error) {
	ep, err := a.getContentIDEndpoint(id)
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addContentQueryParams(query).Encode()
	return a.SendContentRequest(ep, "GET", nil)
}

// ContentSearch results
type ContentSearch struct {
	Results []Content `json:"results"`
	Start   int       `json:"start,omitempty"`
	Limit   int       `json:"limit,omitempty"`
	Size    int       `json:"size,omitempty"`
}

// GetContent querys content using a query parameters
func (a *API) GetContent(query ContentQuery) (*ContentSearch, error) {
	ep, err := a.getContentEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addContentQueryParams(query).Encode()

	req, err := http.NewRequest("GET", ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var search ContentSearch

	err = json.Unmarshal(res, &search)
	if err != nil {
		return nil, err
	}
	return &search, nil
}

// GetChildPages returns a content list of child page objects
func (a *API) GetChildPages(id string) (*Search, error) {
	var (
		results      []Results
		searchResult Search
	)

	ep, err := a.getContentChildEndpoint(id, "page")
	if err != nil {
		return nil, err
	}

	query := ContentQuery{
		Start: 0,
		Limit: 25,
	}

	searchResult.Start = 0

	for {
		ep.RawQuery = addContentQueryParams(query).Encode()
		s, err := a.SendSearchRequest(ep, "GET")
		if err != nil {
			return nil, err
		}
		results = append(results, s.Results...)
		if len(s.Results) < query.Limit {
			break
		}
		query.Start += 25
	}

	searchResult.Limit = len(results)
	searchResult.Size = len(results)
	searchResult.Results = results

	return &searchResult, nil
}

// GetComments returns a list of comments belonging to id
func (a *API) GetComments(id string) (*Search, error) {
	ep, err := a.getContentChildEndpoint(id, "comment")
	if err != nil {
		return nil, err
	}
	return a.SendSearchRequest(ep, "GET")
}

// GetAttachments returns a list of attachments belonging to id
func (a *API) GetAttachments(id string) (*Search, error) {
	ep, err := a.getContentChildEndpoint(id, "attachment")
	if err != nil {
		return nil, err
	}
	return a.SendSearchRequest(ep, "GET")
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

// GetHistory returns history information
func (a *API) GetHistory(id string) (*History, error) {
	ep, err := a.getContentGenericEndpoint(id, "history")
	if err != nil {
		return nil, err
	}
	return a.SendHistoryRequest(ep, "GET")
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

// GetLabels returns a list of labels attachted to a content object
func (a *API) GetLabels(id string) (*Labels, error) {
	ep, err := a.getContentGenericEndpoint(id, "label")
	if err != nil {
		return nil, err
	}
	return a.SendLabelRequest(ep, "GET", nil)
}

// AddLabels returns adds labels
func (a *API) AddLabels(id string, labels *[]Label) (*Labels, error) {
	ep, err := a.getContentGenericEndpoint(id, "label")
	if err != nil {
		return nil, err
	}
	return a.SendLabelRequest(ep, "POST", labels)
}

// DeleteLabel removes a label by name from content identified by id
func (a *API) DeleteLabel(id string, name string) (*Labels, error) {
	ep, err := a.getContentGenericEndpoint(id, "label/"+name)
	if err != nil {
		return nil, err
	}
	return a.SendLabelRequest(ep, "DELETE", nil)
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

// GetWatchers returns a list of watchers
func (a *API) GetWatchers(id string) (*Watchers, error) {
	ep, err := a.getContentGenericEndpoint(id, "notification/child-created")
	if err != nil {
		return nil, err
	}
	return a.SendWatcherRequest(ep, "GET")
}

// CreateContent creates content
func (a *API) CreateContent(c *Content) (*Content, error) {
	ep, err := a.getContentEndpoint()
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "POST", c)
}

// UpdateContent updates content
func (a *API) UpdateContent(c *Content) (*Content, error) {
	ep, err := a.getContentIDEndpoint(c.ID)
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "PUT", c)
}

// UploadAttachment uploaded the given reader as an attachment to the
// page with the given id. The existing attachment won't be updated with
// a new version number
func (a *API) UploadAttachment(id string, attachmentName string, attachment io.Reader) (*Search, error) {
	ep, err := a.getContentChildEndpoint(id, "attachment")
	if err != nil {
		return nil, err
	}
	return a.SendContentAttachmentRequest(ep, attachmentName, attachment, map[string]string{})
}

// UpdateAttachment update the attachment with an attachmentID on a page with an id to a new version
func (a *API) UpdateAttachment(id string, attachmentName string, attachmentID string, attachment io.Reader) (*Search, error) {
	ep, err := a.getContentChildEndpoint(id, "attachment/"+attachmentID+"/data")
	if err != nil {
		return nil, err
	}
	return a.SendContentAttachmentRequest(ep, attachmentName, attachment, map[string]string{})
}

// DelContent deletes content by id
func (a *API) DelContent(id string) (*Content, error) {
	ep, err := a.getContentIDEndpoint(id)
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "DELETE", nil)
}

// ContentVersionResult contains the version results

/*
type ContentVersionResult struct {
	Result []Version `json:"results"`
}
*/

// getContentGenericEndpoint creates the correct api endpoint by given id and type
func (a *API) getContentVersionEndpoint(id string, t string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/experimental/content/" + id + "/" + t)
}

// GetContentVersion gets all versions of this content
func (a *API) GetContentVersion(id string) (*ContentVersionResult, error) {
	ep, err := a.getContentVersionEndpoint(id, "version")
	if err != nil {
		return nil, err
	}
	return a.SendContentVersionRequest(ep, "GET")
}

// addContentQueryParams adds the defined query parameters
func addContentQueryParams(query ContentQuery) *url.Values {

	data := url.Values{}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	//get specific version
	if query.Version != 0 {
		data.Set("version", strconv.Itoa(query.Version))
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.OrderBy != "" {
		data.Set("orderby", query.OrderBy)
	}
	if query.PostingDay != "" {
		data.Set("postingDay", query.PostingDay)
	}
	if query.SpaceKey != "" {
		data.Set("spaceKey", query.SpaceKey)
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	if query.Status != "" {
		data.Set("status", query.Status)
	}
	if query.Title != "" {
		data.Set("title", query.Title)
	}
	if query.Trigger != "" {
		data.Set("trigger", query.Trigger)
	}
	if query.Type != "" {
		data.Set("type", query.Type)
	}
	return &data
}

func (a *API) GetPageId(spacename string, pagename string) (*ContentSearch, error) {

	var query ContentQuery

	query.Start = 0
	query.Limit = 10
	query.SpaceKey = spacename
	query.Title = pagename

	return a.GetContent(query)
}

func (a *API) UppdateAttachment(spacename string, pagename string, filename string) error {

	res, err := a.GetPageId(spacename, pagename)
	if err != nil {
		return err
	}

	if res.Size == 1 {
		file, err3 := os.Open(filename)
		if err3 != nil {
			log.Fatal(err3)
		}

		reader := bufio.NewReader(file)

		found := false
		pageid := res.Results[0].ID
		search, err2 := a.GetAttachments(pageid)
		if err2 != nil {
			log.Fatal(err2)
		}
		_, name := filepath.Split(filename)
		for _, v := range search.Results {
			if v.Title == name {
				_, e := a.UpdateAttachment(pageid, name, v.ID, reader)
				if e != nil {
					return e
				}
				found = true
			}
		}
		if !found {
			_, e := a.UploadAttachment(pageid, name, reader)
			if e != nil {
				return e
			}
		}

	} else {
		return errors.New("page not found")
	}
	return nil
}

// AddPage adds a new page to the space with the given title, TODO what if page already exists?
func (a *API) AddPage(title, spaceKey, filepath string, bodyOnly, stripImgs bool, ancestor string) {

	// create content
	data := &Content{
		Type:  "page",
		Title: title,
		Ancestors: []Ancestor{
			Ancestor{
				ID: ancestor,
			},
		},
		Body: Body{
			Storage: Storage{
				Value:          "",
				Representation: "storage",
			},
		},
		Version: &Version{
			Number: 1,
		},
		Space: Space{
			Key: spaceKey,
		},
	}
	data.Body.Storage.Value = getBodyFromFile(filepath, bodyOnly, stripImgs)

	_, err := a.CreateContent(data)
	if err != nil {
		log.Fatal(err)
	}

}

func getBodyFromFile(filepath string, bodyOnly, stripImgs bool) string {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if !bodyOnly {
		return string(buf)
	}
	return StripHTML(buf, bodyOnly, stripImgs)
}
