package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1_TokenReviewStatus.go


// UserInfo holds the information about the user needed to implement the user.Info interface.
type Io_k8s_api_authentication_v1_UserInfo struct {
	// Any additional information provided by the authenticator.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Extra    *interface{} `json:"extra,omitempty"`

	// The names of groups this user is a part of.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Groups   []*string    `json:"groups,omitempty"`

	// A unique value that identifies this user across time. If this user is deleted and another user by the same name is
	// added, they will have different UIDs.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Uid      *string      `json:"uid,omitempty"`

	// The name that uniquely identifies this user among all active users.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Username *string      `json:"username,omitempty"`
}
