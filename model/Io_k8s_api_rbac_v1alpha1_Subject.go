package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_ClusterRoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_RoleBinding.go


// Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct
// API object reference, or a value for non-objects such as user and group names.
type Io_k8s_api_rbac_v1alpha1_Subject struct {
	// APIVersion holds the API group and version of the referenced subject. Defaults to "v1" for ServiceAccount subjects.
	// Defaults to "rbac.authorization.k8s.io/v1alpha1" for User and Group subjects.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the
	// Authorizer does not recognized the kind value, the Authorizer should report an error.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string `json:"kind"`

	// Name of the object being referenced.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name       *string `json:"name"`

	// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is
	// not empty the Authorizer should report an error.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace  *string `json:"namespace,omitempty"`
}
