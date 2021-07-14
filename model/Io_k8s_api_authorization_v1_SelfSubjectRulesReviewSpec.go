package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SelfSubjectRulesReview.go
type Io_k8s_api_authorization_v1_SelfSubjectRulesReviewSpec struct {
	// Namespace to evaluate rules for. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace *string `json:"namespace,omitempty"`
}
