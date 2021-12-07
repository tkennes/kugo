package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authentication_v1_TokenRequestSpec.go


// BoundObjectReference is a reference to an object that a token is bound to.
type Io_k8s_api_authentication_v1_BoundObjectReference struct {
	// API version of the referent.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string `json:"kind,omitempty"`

	// Name of the referent.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name       *string `json:"name,omitempty"`

	// UID of the referent.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Uid        *string `json:"uid,omitempty"`
}
