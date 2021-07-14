package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_RoleBindingList.go


// RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole
// in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.
// RoleBindings in a given namespace only have effect in that namespace.
type Io_k8s_api_rbac_v1_RoleBinding struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                          `json:"kind,omitempty"`

	// Standard object's metadata.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be
	// resolved, the Authorizer must return an error.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_RoleRef.go
	RoleRef    Io_k8s_api_rbac_v1_RoleRef                       `json:"roleRef"`

	// Subjects holds references to the objects the role applies to.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_Subject.go
	Subjects   []Io_k8s_api_rbac_v1_Subject                     `json:"subjects,omitempty"`
}
