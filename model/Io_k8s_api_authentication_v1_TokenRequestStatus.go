package kugo_model

import (
	"time"
)


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authentication_v1_TokenRequest.go


// TokenRequestStatus is the result of a token request.
type Io_k8s_api_authentication_v1_TokenRequestStatus struct {
	// ExpirationTimestamp is the time of expiration of the returned token.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	ExpirationTimestamp time.Time `json:"expirationTimestamp"`

	// Token is the opaque bearer token.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Token               *string   `json:"token"`
}
