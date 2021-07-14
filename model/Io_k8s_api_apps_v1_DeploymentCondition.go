package model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_DeploymentStatus.go


// DeploymentCondition describes the state of a deployment at a certain point.
type Io_k8s_api_apps_v1_DeploymentCondition struct {
	// Last time the condition transitioned from one status to another.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// The last time this condition was updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastUpdateTime     *time.Time `json:"lastUpdateTime,omitempty"`

	// A human readable message indicating details about the transition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// The reason for the condition's last transition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status             *string    `json:"status"`

	// Type of deployment condition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string    `json:"type"`
}
