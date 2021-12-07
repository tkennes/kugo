package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimStatus.go


// PersistentVolumeClaimCondition contails details about state of pvc
type Io_k8s_api_core_v1_PersistentVolumeClaimCondition struct {
	// Last time we probed the condition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastProbeTime      *time.Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If
	// it reports "ResizeStarted" that means the underlying persistent volume is being resized.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Status             *string    `json:"status"`

	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type               *string    `json:"type"`
}
