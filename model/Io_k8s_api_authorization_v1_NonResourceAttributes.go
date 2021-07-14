package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SubjectAccessReviewSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SelfSubjectAccessReviewSpec.go


// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer
// interface
type Io_k8s_api_authorization_v1_NonResourceAttributes struct {
	// Path is the URL path of the request
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path *string `json:"path,omitempty"`

	// Verb is the standard HTTP verb
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Verb *string `json:"verb,omitempty"`
}
