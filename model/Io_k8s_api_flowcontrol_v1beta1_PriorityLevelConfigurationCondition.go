package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationStatus.go


// PriorityLevelConfigurationCondition defines the condition of priority level.
type Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationCondition struct {
	// `lastTransitionTime` is the last time the condition transitioned from one status to another.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// `message` is a human-readable message indicating details about last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// `status` is the status of the condition. Can be True, False, Unknown. Required.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Status             *string    `json:"status,omitempty"`

	// `type` is the type of the condition. Required.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type               *string    `json:"type,omitempty"`
}
