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

var default_headers = http.Header{
	"Content-Type": {"application/json"},
	"Accept":       {"application/json"},
}

var experimental_headers = http.Header{
	"Content-Type":      {"application/json"},
	"Accept":            {"application/json"},
	"X-ExperimentalApi": {"opt-in"},
}
var form_token_headers = http.Header{
	"Content-Type":      {"application/x-www-form-urlencoded; charset=UTF-8"},
	"X-Atlassian-Token": {"no-check"},
}
var no_check_headers = http.Header{"X-Atlassian-Token": {"no-check"}}
var safe_mode_headers = http.Header{
	"X-Atlassian-Token": {"nocheck"},
	"Content-Type":      {"application/vnd.atl.plugins.safe.mode.flag+json"},
}
var experimental_headers_general = http.Header{
	"X-Atlassian-Token": {"no-check"},
	"X-ExperimentalApi": {"opt-in"},
}

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
	case http.StatusOK, http.StatusCreated, http.StatusPartialContent, 202:
		return res, nil
	case http.StatusNoContent, http.StatusResetContent:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("authentication failed")
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("service is not available: %s", resp.Status)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error: %s", resp.Status)
	case http.StatusForbidden:
		return nil, fmt.Errorf("forbidden: %s", resp.Status)
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

	part, err := writer.CreateFormFile("file", attachmentName)
	if err != nil {
		return nil, err
	}

	// add attachment to body
	_, err = io.Copy(part, attachment)
	if err != nil {
		return nil, err
	}

	// add any other params
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	//clean up multipart form writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ep.String(), body) // will always be put
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Atlassian-Token", "nocheck") // required by api
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// https://developer.atlassian.com/cloud/confluence/rest/#api-api-content-id-child-attachment-put

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

// SendClusterRequest sends cluster related requests
func (a *API) SendClusterRequest(ep *url.URL, method string) (*Cluster, error) {

	req, err := http.NewRequest(method, ep.String(), nil)

	// https://confluence.atlassian.com/cloudkb/xsrf-check-failed-when-calling-cloud-apis-826874382.html
	req.Header = no_check_headers

	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("\n%s\n", res)
	var cluster Cluster

	err = json.Unmarshal(res, &cluster)
	if err != nil {
		return nil, err
	}

	return &cluster, nil
}

// SendHealthCheckStatusesRequest TODO
func (a *API) SendHealthCheckStatusesRequest(ep *url.URL, method string) (*HealthCheckStatuses, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("\n%s\n", res)
	var healthCheckStatuses HealthCheckStatuses

	err = json.Unmarshal(res, &healthCheckStatuses)
	if err != nil {
		return nil, err
	}

	return &healthCheckStatuses, nil
}

// SendPreUpgradeInfoRequest TODO
func (a *API) SendPreUpgradeInfoRequest(ep *url.URL, method string) (*PreUpgradeInfo, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("\n%s\n", res)
	var preUpgradeInfo PreUpgradeInfo

	err = json.Unmarshal(res, &preUpgradeInfo)
	if err != nil {
		return nil, err
	}

	return &preUpgradeInfo, nil
}

func (a *API) SendProductUpdatesRequest(ep *url.URL, method string) (*ProductUpdates, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("\n%s\n", res)
	var pluginUpdates ProductUpdates

	err = json.Unmarshal(res, &pluginUpdates)
	if err != nil {
		return nil, err
	}

	return &pluginUpdates, nil
}

func (a *API) SendProductUpdateCompatibilitiesRequest(ep *url.URL, method string) (*ProductUpdateCompatibilities, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("\n%s\n", res)
	var productUpdateCompatibilities ProductUpdateCompatibilities

	err = json.Unmarshal(res, &productUpdateCompatibilities)
	if err != nil {
		return nil, err
	}

	return &productUpdateCompatibilities, nil
}

func (a *API) SendPluginMarketplaceInfosRequest(ep *url.URL, method string) (*PluginMarketplaceInfos, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var pluginMarketplaceInfos PluginMarketplaceInfos

	err = json.Unmarshal(res, &pluginMarketplaceInfos)
	if err != nil {
		return nil, err
	}

	return &pluginMarketplaceInfos, nil
}

func (a *API) SendUpmTokenRequest(ep *url.URL, method string) (*string, error) {

	req, err := http.NewRequest(method, ep.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.atl.plugins.installed+json")

	if (a.username != "") || (a.token != "") {
		a.Auth(req)
	}

	res, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var upm string
	upm = res.Header.Get("Upm-Token")
	if err != nil {
		return nil, err
	}

	return &upm, nil
}

func (a *API) SendPluginUpdateRequest(ep *url.URL, method, pluginBinaryUri, pluginName, pluginVersion string) (http.Header, error) {

	var jsonStr = []byte(fmt.Sprintf(`{"pluginUri":"%s","pluginName": "%s","pluginVersion": "%s"}`, pluginBinaryUri, pluginName, pluginVersion))
	req, err := http.NewRequest(method, ep.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/vnd.atl.plugins.install.uri+json")
	req.Header.Set("Accept", "application/json")

	if (a.username != "") || (a.token != "") {
		a.Auth(req)
	}

	res, err := a.Client.Do(req)
	fmt.Printf("\n%+v\n", res)
	if err != nil {
		return nil, err
	}

	return res.Header, nil
}
