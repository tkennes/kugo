package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_Subject.go


// GroupSubject holds detailed information for group-kind subject.
type Io_k8s_api_flowcontrol_v1alpha1_GroupSubject struct {
	// name is the user group that matches, or "*" to match all user groups. See
	// https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names.
	// Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name *string `json:"name"`
}
