package goconfluence

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Request implements the basic Request function
func (a *API) Request(req *http.Request) ([]byte, error) {
	req.Header.Add("Accept", "application/json, */*")
	a.Auth(req)

	Debug("====== Request ======")
	Debug(req)
	Debug("====== /Request ======")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	Debug(fmt.Sprintf("====== Response Status Code: %d ======", resp.StatusCode))

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	Debug("====== Response Body ======")
	Debug(string(res))
	Debug("====== /Response Body ======")

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusPartialContent:
		return res, nil
	case http.StatusNoContent, http.StatusResetContent:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("authentication failed")
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("service is not available: %s", resp.Status)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error: %s", resp.Status)
	case http.StatusConflict:
		return nil, fmt.Errorf("conflict: %s", resp.Status)
	}

	return nil, fmt.Errorf("unknown response status: %s", resp.Status)
}

// SendContentRequest sends content related requests
// this function is used for getting, updating and deleting content
func (a *API) SendContentRequest(ep *url.URL, method string, c *Content) (*Content, error) {

	var body io.Reader
	if c != nil {
		js, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(js))
	}

	req, err := http.NewRequest(method, ep.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var content Content

	err = json.Unmarshal(res, &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

// SendUserRequest sends user related requests
func (a *API) SendUserRequest(ep *url.URL, method string) (*User, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var user User

	err = json.Unmarshal(res, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SendSearchRequest sends search related requests
func (a *API) SendSearchRequest(ep *url.URL, method string) (*Search, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var search Search

	err = json.Unmarshal(res, &search)
	if err != nil {
		return nil, err
	}

	return &search, nil
}

// SendHistoryRequest requests history
func (a *API) SendHistoryRequest(ep *url.URL, method string) (*History, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var history History

	err = json.Unmarshal(res, &history)
	if err != nil {
		return nil, err
	}

	return &history, nil
}

// SendLabelRequest requests history
func (a *API) SendLabelRequest(ep *url.URL, method string, labels *[]Label) (*Labels, error) {

	var body io.Reader
	if labels != nil {
		js, err := json.Marshal(labels)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(js))
	}

	req, err := http.NewRequest(method, ep.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	if res != nil {
		var l Labels

		err = json.Unmarshal(res, &l)
		if err != nil {
			return nil, err
		}

		return &l, nil
	}

	return &Labels{}, nil
}

// SendWatcherRequest requests watchers
func (a *API) SendWatcherRequest(ep *url.URL, method string) (*Watchers, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}
	var watchers Watchers

	err = json.Unmarshal(res, &watchers)
	if err != nil {
		return nil, err
	}

	return &watchers, nil
}

// SendAllSpacesRequest sends a request for all spaces
func (a *API) SendAllSpacesRequest(ep *url.URL, method string) (*AllSpaces, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var allSpaces AllSpaces

	err = json.Unmarshal(res, &allSpaces)
	if err != nil {
		return nil, err
	}

	return &allSpaces, nil
}
