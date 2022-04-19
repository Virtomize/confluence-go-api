package goconfluence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Request implements the basic Request function
func (a *API) Request(req *http.Request) ([]byte, error) {
	req.Header.Add("Accept", "application/json, */*")

	// only auth if we can auth
	if (a.username != "") || (a.token != "") {
		a.Auth(req)
	}

	Debug("====== Request ======")
	Debug(req)
	Debug("====== Request Body ======")
	if DebugFlag {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
	}
	Debug("====== /Request Body ======")
	Debug("====== /Request ======")

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	Debug(fmt.Sprintf("====== Response Status Code: %d ======", resp.StatusCode))

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

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
	if len(res) != 0 {
		err = json.Unmarshal(res, &content)
		if err != nil {
			return nil, err
		}
	}
	return &content, nil
}

// SendContentAttachmentRequest sends a multipart/form-data attachment create/update request to a content
func (a *API) SendContentAttachmentRequest(ep *url.URL, attachmentName string, attachment io.Reader, params map[string]string) (*Search, error) {
	// setup body for mulitpart file, adding minorEdit option
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := writer.WriteField("minorEdit", "true")
	if err != nil {
		return nil, err
	}
	part, err2 := writer.CreateFormFile("file", attachmentName)
	if err2 != nil {
		return nil, err2
	}

	// add attachment to body
	_, err2 = io.Copy(part, attachment)
	if err2 != nil {
		return nil, err2
	}

	// add any other params
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	//clean up multipart form writer
	err2 = writer.Close()
	if err2 != nil {
		return nil, err2
	}

	req, err2 := http.NewRequest("POST", ep.String(), body) // will always be put
	if err2 != nil {
		return nil, err2
	}

	req.Header.Set("X-Atlassian-Token", "nocheck") // required by api
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// https://developer.atlassian.com/cloud/confluence/rest/#api-api-content-id-child-attachment-put

	res, err2 := a.Request(req)
	if err2 != nil {
		return nil, err2
	}

	var search Search

	err2 = json.Unmarshal(res, &search)
	if err2 != nil {
		return nil, err2
	}

	return &search, nil
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

	if a.Debug {
		fmt.Printf("Send: %s, Method: %s \n", ep.String(), method)
	}
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
	if a.Debug {
		fmt.Printf("Reply: %s \n", res)
	}

	return &allSpaces, nil
}

// SendContentVersionRequest requests a version of a specific content
func (a *API) SendContentVersionRequest(ep *url.URL, method string) (*ContentVersionResult, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var versionResult ContentVersionResult

	err = json.Unmarshal(res, &versionResult)
	if err != nil {
		return nil, err
	}

	return &versionResult, nil
}

func (a *API) getSendGenericRequest(ep string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + ep)
}

// SendGenericRequest sends a greneric request
func (a *API) SendGenericRequest(ep, method string) ([]byte, error) {
	endp, err := a.getSendGenericRequest(ep)
	if err != nil {
		return nil, err
	}

	if a.Debug {
		fmt.Printf("Send: %s, Method: %s \n", endp, method)
	}
	req, err := http.NewRequest(method, endp.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
