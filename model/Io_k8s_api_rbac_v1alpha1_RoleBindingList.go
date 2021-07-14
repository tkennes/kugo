package model


// RoleBindingList is a collection of RoleBindings Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1
// RoleBindingList, and will no longer be served in v1.22.
type Io_k8s_api_rbac_v1alpha1_RoleBindingList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                        `json:"apiVersion,omitempty"`

	// Items is a list of RoleBindings
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_RoleBinding.go
	Items      []Io_k8s_api_rbac_v1alpha1_RoleBinding         `json:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                        `json:"kind,omitempty"`

	// Standard object's metadata.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta `json:"metadata,omitempty"`
}
