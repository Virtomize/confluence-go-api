package goconfluence

import (
	"fmt"
	"net/url"
)

type ProductUpdates struct {
	Versions   []ProductUpdateConfluenceVersion `json:"versions,omitempty"`
}

type ProductUpdateLink struct {
	Self string `json:"self,omitempty"`
}

type ProductUpdateConfluenceVersion struct {
	Version string `json:"version,omitempty"`
	Recent  bool   `json:"recent,omitempty"`
	Links   ProductUpdateLink  `json:"links,omitempty"`
}

type ProductUpdateCompatibilities struct {
	Compatible                       []ProductUpdatePluginCompatibility `json:"compatible,omitempty"`
	UpdateRequired                   []ProductUpdatePluginCompatibility `json:"updateRequired,omitempty"`
	UpdateRequiredAfterProductUpdate []ProductUpdatePluginCompatibility `json:"updateRequiredAfterProductUpdate,omitempty"`
	Incompatible                     []ProductUpdatePluginCompatibility `json:"incompatible,omitempty"`
	Unknown                          []ProductUpdatePluginCompatibility `json:"unknown,omitempty"`
}

type ProductUpdatePluginCompatibility struct {
	Links   ProductUpdatePluginCompatibilityLink  `json:"links,omitempty"`
	Name    string `json:"name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Key     string `json:"key,omitempty"`
}

type ProductUpdatePluginCompatibilityLink struct {
	Modify     string `json:"modify,omitempty"`
	Self       string `json:"self,omitempty"`
	PluginIcon string `json:"plugin-icon,omitempty"`
	PluginLogo string `json:"plugin-logo,omitempty"`
}

type PluginMarketplaceInfos struct {
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Update Update `json:"update,omitempty"`
}

type Update struct {
	Links       Link   `json:"links,omitempty"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
}

type Link struct {
	Binary string `json:"binary,omitempty"`
}

func (a *API) PluginUpdates() (*ProductUpdates, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/plugins/1.0/product-updates/")
	if err != nil {
		return nil, err
	}
	return a.SendProductUpdatesRequest(ep, "GET")
}

func (a *API) PluginUpdateCompatibility(compatibilityLink string) (*ProductUpdateCompatibilities, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + compatibilityLink)
        fmt.Printf("\n%s\n", ep)
	if err != nil {
		return nil, err
	}
	return a.SendProductUpdateCompatibilitiesRequest(ep, "GET")
}

func (a *API) PluginMarketplaceInfos(pluginKey string) (*PluginMarketplaceInfos, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/plugins/1.0/" + pluginKey + "/marketplace")
	fmt.Printf("\n%s\n", ep)
	if err != nil {
		return nil, err
	}
	return a.SendPluginMarketplaceInfosRequest(ep, "GET")
}

func (a *API) GetUpmToken() {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/plugins/1.0/?os_authType=basic")
	// accept: application/vnd.atl.plugins.installed+json
	fmt.Printf("\n%s\n", ep)
	if err != nil {
		fmt.Print(err)
	}
	a.SendGetUpmToken(ep, "GET")
}
