package goconfluence

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// getContentIDEndpoint creates the correct api endpoint by given id
func (a *API) getContentIDEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/content/" + id)
}

// getContentEndpoint creates the correct api endpoint
func (a *API) getContentEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/content/")
}

// getContentChildEndpoint creates the correct api endpoint by given id and type
func (a *API) getContentChildEndpoint(id string, t string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/content/" + id + "/child/" + t)
}

// getContentGenericEndpoint creates the correct api endpoint by given id and type
func (a *API) getContentGenericEndpoint(id string, t string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/content/" + id + "/" + t)
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

// GetHistory returns history information
func (a *API) GetHistory(id string) (*History, error) {
	ep, err := a.getContentGenericEndpoint(id, "history")
	if err != nil {
		return nil, err
	}
	return a.SendHistoryRequest(ep, "GET")
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

// UpdateAttachment update the attachment with an attachmentId on a page with an id to a new version
func (a *API) UpdateAttachment(id string, attachmentName string, attachmentId string, attachment io.Reader) (*Search, error) {
	ep, err := a.getContentChildEndpoint(id, "attachment/"+attachmentId+"/data")
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
