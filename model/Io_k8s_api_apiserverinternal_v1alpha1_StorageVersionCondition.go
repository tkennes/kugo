package model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionStatus.go


// Describes the state of the storageVersion at a certain point.
type Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionCondition struct {
	// Last time the condition transitioned from one status to another.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// A human readable message indicating details about the transition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// If set, this represents the .metadata.generation that the condition was set based upon.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration *int       `json:"observedGeneration,omitempty"`

	// The reason for the condition's last transition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason             *string    `json:"reason"`

	// Status of the condition, one of True, False, Unknown.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status             *string    `json:"status"`

	// Type of the condition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string    `json:"type"`
}
