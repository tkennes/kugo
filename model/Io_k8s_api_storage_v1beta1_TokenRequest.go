package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_CSIDriverSpec.go


// TokenRequest contains parameters of a service account token.
type Io_k8s_api_storage_v1beta1_TokenRequest struct {
	// Audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube
	// apiserver.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Audience          *string `json:"audience"`

	// ExpirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of
	// "ExpirationSeconds" in "TokenRequestSpec"
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	ExpirationSeconds *int    `json:"expirationSeconds,omitempty"`
}
