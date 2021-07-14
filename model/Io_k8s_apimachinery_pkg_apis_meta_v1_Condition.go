package model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceStatus.go


// Condition contains details for one aspect of the current state of this API Resource.
type Io_k8s_apimachinery_pkg_apis_meta_v1_Condition struct {
	// lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the
	// underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastTransitionTime time.Time `json:"lastTransitionTime"`

	// message is a human readable message indicating details about the transition. This may be an empty string.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message            *string   `json:"message"`

	// observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if
	// .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of
	// date with respect to the current state of the instance.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration *int      `json:"observedGeneration,omitempty"`

	// reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of
	// specific condition types may define expected values and meanings for this field, and whether the values are considered a
	// guaranteed API. The value should be a CamelCase string. This field may not be empty.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason             *string   `json:"reason"`

	// status of the condition, one of True, False, Unknown.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status             *string   `json:"status"`

	// type of condition in CamelCase or in foo.example.com/CamelCase.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string   `json:"type"`
}
