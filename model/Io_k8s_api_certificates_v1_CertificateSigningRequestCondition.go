package model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1_CertificateSigningRequestStatus.go


// CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object
type Io_k8s_api_certificates_v1_CertificateSigningRequestCondition struct {
	// lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new
	// condition type is added or an existing condition's status is changed, the server defaults this to the current time.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`

	// lastUpdateTime is the time of the last update to this condition
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastUpdateTime     *time.Time `json:"lastUpdateTime,omitempty"`

	// message contains a human readable message with details about the request state
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message            *string    `json:"message,omitempty"`

	// reason indicates a brief reason for the request state
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason             *string    `json:"reason,omitempty"`

	// status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or
	// "Unknown".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status             *string    `json:"status"`

	// type of the condition. Known conditions are "Approved", "Denied", and "Failed".  An "Approved" condition is added via
	// the /approval subresource, indicating the request was approved and should be issued by the signer.  A "Denied" condition
	// is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.  A
	// "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.
	// Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once
	// added.  Only one condition of a given type is allowed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string    `json:"type"`
}
