package goconfluence

import (
	"net/url"
)

// https://confluence-test.dev.itops.breuni.de/rest/zdu/cluster
type HealthCheckStatuses struct {
	Statuses []HealthCheckStatus `json:"statuses"`
}

type HealthCheckStatus struct {
	ID              int    `json:"id"`
	IsSoftLaunch    bool   `json:"isSoftLaunch"`
	IsEnabled       bool   `json:"isEnabled"`
	CompleteKey     string `json:"completeKey"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	IsHealthy       bool   `json:"isHealthy"`
	FailureReason   string `json:"failureReason"`
	Application     string `json:"application"`
	NodeID          string `json:"nodeId,omitempty"`
	Time            int64  `json:"time"`
	Severity        string `json:"severity"`
	Documentation   string `json:"documentation"`
	Tag             string `json:"tag"`
	AdditionalLinks []any  `json:"additionalLinks"`
	Enabled         bool   `json:"enabled"`
	Healthy         bool   `json:"healthy"`
	SoftLaunch      bool   `json:"softLaunch"`
}

type PreUpgradeInfo struct {
	Versions        []ConfluenceVersion `json:"versions,omitempty"`
	SelectedVersion Version             `json:"selectedVersion,omitempty"`
	InstanceData    InstanceData        `json:"instanceData,omitempty"`
	Stale           bool                `json:"stale,omitempty"`
}

type SubSection struct {
	Description string `json:"description,omitempty"`
	Steps       []any  `json:"steps,omitempty"`
}

type UpgradePath struct {
	AnalyticsKey string       `json:"analyticsKey,omitempty"`
	Title        string       `json:"title,omitempty"`
	SubSections  []SubSection `json:"subSections,omitempty"`
}

type ConfluenceVersion struct {
	FullName               string        `json:"fullName,omitempty"`
	ShortName              string        `json:"shortName,omitempty"`
	AnalyticsVersion       string        `json:"analyticsVersion,omitempty"`
	UpgradeInstructionsURL string        `json:"upgradeInstructionsUrl,omitempty"`
	InstallerURL           string        `json:"installerUrl,omitempty"`
	ArchiveURL             string        `json:"archiveUrl,omitempty"`
	ReleaseDate            int64         `json:"releaseDate,omitempty"`
	ReleaseNotesURL        string        `json:"releaseNotesUrl,omitempty"`
	SupportedPlatforms     []any         `json:"supportedPlatforms,omitempty"`
	UpgradePath            []UpgradePath `json:"upgradePath,omitempty"`
	ModifiedFiles          []string      `json:"modifiedFiles,omitempty"`
	IsZduAvailable         bool          `json:"isZduAvailable,omitempty"`
}

type InstanceData struct {
	PlatformID         string `json:"platformId,omitempty"`
	FullName           string `json:"fullName,omitempty"`
	ProductDisplayName string `json:"productDisplayName,omitempty"`
	UpgradeURL         string `json:"upgradeUrl,omitempty"`
	ReleaseDate        int64  `json:"releaseDate,omitempty"`
	AnalyticsVersion   string `json:"analyticsVersion,omitempty"`
	OperatingSystem    string `json:"operatingSystem,omitempty"`
}

func (a *API) HealthCheckStatuses() (*HealthCheckStatuses, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/troubleshooting/1.0/check")
	if err != nil {
		return nil, err
	}
	return a.SendHealthCheckStatusesRequest(ep, "GET")
}

func (a *API) PreUpgradeInfo() (*PreUpgradeInfo, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/rest/troubleshooting/latest/pre-upgrade/info")
	if err != nil {
		return nil, err
	}
	return a.SendPreUpgradeInfoRequest(ep, "GET")
}
