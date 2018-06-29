package goconfluence

import (
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
func (a *API) GetContentByID(id string) (*Content, error) {
	ep, err := a.getContentIDEndpoint(id)
	if err != nil {
		return nil, err
	}
	return a.SendContentRequest(ep, "GET", nil)
}

// GetContent querys content using a query parameters
func (a *API) GetContent(query ContentQuery) (*Search, error) {
	ep, err := a.getContentEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addContentQueryParams(query).Encode()
	return a.SendSearchRequest(ep, "GET")
}

// GetChildPages returns a content list of child page objects
func (a *API) GetChildPages(id string) (*Search, error) {
	ep, err := a.getContentChildEndpoint(id, "page")
	if err != nil {
		return nil, err
	}
	return a.SendSearchRequest(ep, "GET")
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

// GetLabels returns a list of labels attachted to a content object
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
