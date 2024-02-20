package goconfluence

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ProductUpdates struct {
	Versions []ProductUpdateConfluenceVersion `json:"versions,omitempty"`
}

type ProductUpdateLink struct {
	Self string `json:"self,omitempty"`
}

type ProductUpdateConfluenceVersion struct {
	Version string            `json:"version,omitempty"`
	Recent  bool              `json:"recent,omitempty"`
	Links   ProductUpdateLink `json:"links,omitempty"`
}

type ProductUpdateCompatibilities struct {
	Compatible                       []ProductUpdatePluginCompatibility `json:"compatible,omitempty"`
	UpdateRequired                   []ProductUpdatePluginCompatibility `json:"updateRequired,omitempty"`
	UpdateRequiredAfterProductUpdate []ProductUpdatePluginCompatibility `json:"updateRequiredAfterProductUpdate,omitempty"`
	Incompatible                     []ProductUpdatePluginCompatibility `json:"incompatible,omitempty"`
	Unknown                          []ProductUpdatePluginCompatibility `json:"unknown,omitempty"`
}

type ProductUpdatePluginCompatibility struct {
	Links   ProductUpdatePluginCompatibilityLink `json:"links,omitempty"`
	Name    string                               `json:"name,omitempty"`
	Enabled bool                                 `json:"enabled,omitempty"`
	Key     string                               `json:"key,omitempty"`
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

type Link struct {
	Binary string `json:"binary,omitempty"`
}

type PluginInstallationStatusResponse struct {
	Type      string                                 `json:"type,omitempty"`
	PingAfter int                                    `json:"pingAfter,omitempty"`
	Status    PluginInstallationStatusResponseStatus `json:"status,omitempty"`
	Links     PluginInstallationStatusResponseLinks  `json:"links,omitempty"`
	Timestamp int64                                  `json:"timestamp,omitempty"`
	UserKey   string                                 `json:"userKey,omitempty"`
	ID        string                                 `json:"id,omitempty"`
}
type PluginInstallationStatusResponseStatus struct {
	Done        *bool  `json:"done,omitempty"`
	StatusCode  int    `json:"statusCode,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Source      string `json:"source,omitempty"`
	Name        string `json:"name,omitempty"`
}
type PluginInstallationStatusResponseLinks struct {
	Self      string `json:"self,omitempty"`
	Alternate string `json:"alternate,omitempty"`
}

// start
type InstalledMarketplacePlugins struct {
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	Key             string        `json:"key"`
	Name            string        `json:"name"`
	Incompatible    bool          `json:"incompatible"`
	UpdateAvailable bool          `json:"updateAvailable"`
	PrimaryAction   PrimaryAction `json:"primaryAction,omitempty"`
	Update          Update        `json:"update,omitempty"`
}

type PrimaryAction struct {
	Name                            string `json:"name"`
	Priority                        int    `json:"priority"`
	ActionRequired                  bool   `json:"actionRequired"`
	Incompatible                    bool   `json:"incompatible"`
	NonDataCenterApproved           bool   `json:"nonDataCenterApproved"`
	LicenseIncompatibleInDataCenter bool   `json:"licenseIncompatibleInDataCenter"`
}

type Update struct {
	Links                      Link          `json:"links,omitempty"`
	Version                    string        `json:"version,omitempty"`
	VersionDetails             VersionDetail `json:"versionDetails,omitempty"`
	Installable                bool          `json:"installable,omitempty"`
	LicenseCompatible          bool          `json:"licenseCompatible,omitempty"`
	StatusDataCenterCompatible bool          `json:"statusDataCenterCompatible,omitempty"`
}

type VersionDetail struct {
	Deployable bool `json:"deployable"`
	Stable     bool `json:"stable"`
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
	// fmt.Printf("\n%s\n", ep)
	if err != nil {
		return nil, err
	}
	return a.SendProductUpdateCompatibilitiesRequest(ep, "GET")
}

func (a *API) PluginMarketplaceInfos(pluginKey string) (*PluginMarketplaceInfos, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/plugins/1.0/" + pluginKey + "/marketplace")
	// fmt.Printf("\n%s\n", ep)
	if err != nil {
		return nil, err
	}
	return a.SendPluginMarketplaceInfosRequest(ep, "GET")
}

func (a *API) GetUpmToken() (*string, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/plugins/1.0/?os_authType=basic")
	// accept: application/vnd.atl.plugins.installed+json
	// fmt.Printf("\n%s\n", ep)
	if err != nil {
		fmt.Print(err)
	}
	return a.SendUpmTokenRequest(ep, "GET")
}

func (a *API) GetInstalledMarketplacePlugins() (*InstalledMarketplacePlugins, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/plugins/1.0/installed-marketplace?updates=true")
	// accept: application/vnd.atl.plugins.installed+json
	// fmt.Printf("\n%s\n", ep)
	if err != nil {
		fmt.Print(err)
	}
	return a.SendGetInstalledMarketplacePluginsRequest(ep, "GET")
}

func (a *API) UpdatePlugin(pluginBinaryUri, pluginName, pluginVersion, upmToken string) (bool, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + fmt.Sprintf("/rest/plugins/1.0/?token=%s", upmToken))
	if err != nil {
		fmt.Print(err)
	}
	responseHeaders, err := a.SendPluginUpdateRequest(ep, "POST", pluginBinaryUri, pluginName, pluginVersion)
	if err != nil {
		fmt.Print(err)
	}

	// The location header contains the url to the update status check
	updateStatusUrl := responseHeaders.Get("Location")
	if updateStatusUrl == "" {
		return false, errors.New("location header empty")
	}
	// fmt.Printf("\n%s\n", updateStatusUrl)
	ep, err = url.ParseRequestURI(updateStatusUrl)
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)
		req, err := http.NewRequest("GET", ep.String(), nil)
		if err != nil {
			return false, err
		}
		if (a.username != "") || (a.token != "") {
			a.Auth(req)
		}
		res, err := a.Client.Do(req)
		if err != nil {
			fmt.Print("ERROR")
			fmt.Printf("\n%#v\n", res)
			fmt.Print(err)
		}
		fmt.Printf("\n%v\n", res.StatusCode)
		if res.StatusCode != 200 {
			fmt.Print("HTTP ERROR (not 200)")
			fmt.Printf("\n%#v\n", res)
			return false, errors.New("http error blabla")
		}
		contentTypeHeader := string(res.Header.Get("Content-Type"))

		if contentTypeHeader == "application/vnd.atl.plugins.plugin+json" {
			fmt.Print("plugin update done\n")
			return true, nil
		}
		if contentTypeHeader == "application/vnd.atl.plugins.install.installing+json" {
			fmt.Print("plugin update is installing\n")
		}
		if contentTypeHeader == "application/vnd.atl.plugins.install.downloading+json" {
			fmt.Print("plugin update is downloading\n")
		}
	}
	// TODO check if update is actually installed
	return false, errors.New("plugin update did not become ready after X tries")
}
