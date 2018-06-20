/*
	Go library for attlassians confluence wiki

	Copyright (C) 2017 Carsten Seeger

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	@author Carsten Seeger
	@copyright Copyright (C) 2017 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/confluence-go-api
*/

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
	Storage Storage `json:"storage"`
}

// Storage defines the storage information
type Storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

// Version defines the content version number
// the version number is used for updating content
type Version struct {
	Number int `json:"number"`
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
// used for searching
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
}

// User defines user informations
type User struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	UserKey     string `json:"userKey"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}
