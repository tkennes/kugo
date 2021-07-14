package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1_TokenRequest.go


// TokenRequestSpec contains client provided parameters of a token request.
type Io_k8s_api_authentication_v1_TokenRequestSpec struct {
	// Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in
	// the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be
	// used to authenticate against any of the audiences listed but implies a high degree of trust between the target
	// audiences.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Audiences         []*string                                          `json:"audiences"`

	// BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as
	// the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other
	// audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1_BoundObjectReference.go
	BoundObjectRef    *Io_k8s_api_authentication_v1_BoundObjectReference `json:"boundObjectRef,omitempty"`

	// ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a
	// different validity duration so a client needs to check the 'expiration' field in a response.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ExpirationSeconds *int                                               `json:"expirationSeconds,omitempty"`
}
