package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressBackend.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressClassSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressBackend.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressClassSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_IngressBackend.go


// TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same
// namespace.
type Io_k8s_api_core_v1_TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the
	// core API group. For any other third-party types, APIGroup is required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiGroup *string `json:"apiGroup,omitempty"`

	// Kind is the type of resource being referenced
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind     *string `json:"kind"`

	// Name is the name of resource being referenced
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name     *string `json:"name"`
}
