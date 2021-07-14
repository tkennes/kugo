package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1beta1_ClusterRoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1beta1_RoleBinding.go


// RoleRef contains information that points to the role being used
type Io_k8s_api_rbac_v1beta1_RoleRef struct {
	// APIGroup is the group for the resource being referenced
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiGroup *string `json:"apiGroup"`

	// Kind is the type of resource being referenced
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind     *string `json:"kind"`

	// Name is the name of resource being referenced
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name     *string `json:"name"`
}
