package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authentication_v1beta1_TokenReview.go


// TokenReviewSpec is a description of the token authentication request.
type Io_k8s_api_authentication_v1beta1_TokenReviewSpec struct {
	// Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware
	// token authenticators will verify that the token was intended for at least one of the audiences in this list. If no
	// audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Audiences []*string `json:"audiences,omitempty"`

	// Token is the opaque bearer token.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Token     *string   `json:"token,omitempty"`
}
