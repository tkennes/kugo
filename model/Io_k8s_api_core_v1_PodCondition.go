package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodStatus.go


// PodCondition contains details for the current condition of this pod.
type Io_k8s_api_core_v1_PodCondition struct {
	// Last time we probed the condition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastProbeTime      *time.Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Status             *string    `json:"status"`

	// Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-
	// conditions
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type               *string    `json:"type"`
}
