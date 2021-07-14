package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_Subject.go


// UserSubject holds detailed information for user-kind subject.
type Io_k8s_api_flowcontrol_v1alpha1_UserSubject struct {
	// `name` is the username that matches, or "*" to match all usernames. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name *string `json:"name"`
}
