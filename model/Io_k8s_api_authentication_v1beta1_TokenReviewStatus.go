package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1beta1_TokenReview.go


// TokenReviewStatus is the result of the token authentication request.
type Io_k8s_api_authentication_v1beta1_TokenReviewStatus struct {
	// Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token.
	// An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client
	// of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is
	// returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns
	// an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the
	// Kubernetes API server.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Audiences     []*string                                   `json:"audiences,omitempty"`

	// Authenticated indicates that the token was associated with a known user.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Authenticated *bool                                       `json:"authenticated,omitempty"`

	// Error indicates that the token couldn't be checked
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Error         *string                                     `json:"error,omitempty"`

	// User is the UserInfo associated with the provided token.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1beta1_UserInfo.go
	User          *Io_k8s_api_authentication_v1beta1_UserInfo `json:"user,omitempty"`
}
