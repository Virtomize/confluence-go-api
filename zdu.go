package goconfluence

import (
	"net/url"
)

// https://confluence-test.dev.itops.breuni.de/rest/zdu/cluster
type Cluster struct {

	// has finalization tasks
	// Required: true
	HasFinalizationTasks *bool `json:"hasFinalizationTasks"`

	// links
	//Links []*Link `json:"links"`

	// nodes
	Nodes []*ClusterNode `json:"nodes"`

	// original version
	OriginalVersion string `json:"originalVersion,omitempty"`

	// state
	// Enum: [STABLE READY_TO_UPGRADE MIXED READY_TO_RUN_UPGRADE_TASKS RUNNING_UPGRADE_TASKS UPGRADE_TASKS_FAILED]
	State string `json:"state,omitempty"`

	// upgrade mode enabled
	// Required: true
	UpgradeModeEnabled *bool `json:"upgradeModeEnabled"`
}


type ClusterNode struct {

	// active user count
	ActiveUserCount int64 `json:"activeUserCount,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// finalization
	Finalization *ClusterNodeFinalization `json:"finalization,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// links
//	Links []*Link `json:"links"`

	// local
	// Required: true
	Local *bool `json:"local"`

	// name
	Name string `json:"name,omitempty"`

	// state
	// Enum: [STARTING ACTIVE DRAINING TERMINATING OFFLINE ERROR RUNNING_FINALIZE_UPGRADE_TASKS UPGRADE_TASKS_FAILED]
	State string `json:"state,omitempty"`

	// tasks total
	TasksTotal int64 `json:"tasksTotal,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

type ClusterNodeFinalization struct {

	// errors
	Errors []*ClusterNodeFinalizationError `json:"errors"`

	// last requested
	LastRequested int64 `json:"lastRequested,omitempty"`
}

type ClusterNodeFinalizationError struct {

	// cluster upgrade task
	// Required: true
	ClusterUpgradeTask *bool `json:"clusterUpgradeTask"`

	// errors
	Errors []string `json:"errors"`

	// exception message
	ExceptionMessage string `json:"exceptionMessage,omitempty"`
}


// getZduClusterEndpoint creates the correct api endpoint to get the zdu cluster
func (a *API) getZduClusterEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/zdu/cluster")
}

func (a *API) ZduCluster() (*Cluster, error) {
	ep, err := a.getZduClusterEndpoint()
	if err != nil {
		return nil, err
	}
	return a.SendClusterRequest(ep, "GET")
}
