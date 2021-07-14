package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyList.go


// NetworkPolicy describes what network traffic is allowed for a set of Pods
type Io_k8s_api_networking_v1_NetworkPolicy struct {
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

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior for this NetworkPolicy.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_NetworkPolicySpec.go
	Spec       *Io_k8s_api_networking_v1_NetworkPolicySpec      `json:"spec,omitempty"`
}
