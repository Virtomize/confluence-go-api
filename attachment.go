package goconfluence

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Attachment struct {
	MediaTypeDescription string      `json:"mediaTypeDescription"`
	WebuiLink            string      `json:"webuiLink"`
	DownloadLink         string      `json:"downloadLink"`
	CreatedAt            interface{} `json:"createdAt"`
	ID                   string      `json:"id"`
	Comment              string      `json:"comment"`
	Version              struct {
		Number    int       `json:"number"`
		Message   string    `json:"message"`
		MinorEdit bool      `json:"minorEdit"`
		AuthorID  string    `json:"authorId"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"version"`
	Title     string `json:"title"`
	FileSize  int    `json:"fileSize"`
	Status    string `json:"status"`
	PageID    string `json:"pageId"`
	FileID    string `json:"fileId"`
	MediaType string `json:"mediaType"`
	Links     struct {
		Download string `json:"download"`
		Webui    string `json:"webui"`
	} `json:"_links"`
}

// getUserEndpoint creates the correct api endpoint by given id
func (a *API) getAttachmentEndpoint(id string) (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/attachments/" + id)
}

func (a *API) GetAttachmentById(id string) (*Attachment, error) {
	ep, err := a.getAttachmentEndpoint(id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var attachment Attachment

	err = json.Unmarshal(res, &attachment)
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}
