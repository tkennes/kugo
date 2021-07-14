package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1beta1_CertificateSigningRequestStatus.go
type Io_k8s_api_certificates_v1beta1_CertificateSigningRequestCondition struct {
	// lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new
	// condition type is added or an existing condition's status is changed, the server defaults this to the current time.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// timestamp for the last update to this condition
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastUpdateTime     *time.Time `json:"lastUpdateTime,omitempty"`

	// human readable message with details about the request state
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// brief reason for the request state
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or
	// "Unknown". Defaults to "True". If unset, should be treated as "True".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status             *string    `json:"status,omitempty"`

	// type of the condition. Known conditions include "Approved", "Denied", and "Failed".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string    `json:"type"`
}
