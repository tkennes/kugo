package model

import (
	"time"
)


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerState.go


// ContainerStateRunning is a running state of a container.
type Io_k8s_api_core_v1_ContainerStateRunning struct {
	// Time at which the container was last (re-)started
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	StartedAt *time.Time `json:"startedAt,omitempty"`
}
