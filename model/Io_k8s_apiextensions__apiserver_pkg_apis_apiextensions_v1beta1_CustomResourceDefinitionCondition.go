package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionStatus.go


// CustomResourceDefinitionCondition contains details for the current condition of this pod.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionCondition struct {
	// lastTransitionTime last time the condition transitioned from one status to another.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// message is a human-readable message indicating details about last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// reason is a unique, one-word, CamelCase reason for the condition's last transition.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// status is the status of the condition. Can be True, False, Unknown.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Status             *string    `json:"status"`

	// type is the type of the condition. Types include Established, NamesAccepted and Terminating.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type               *string    `json:"type"`
}
