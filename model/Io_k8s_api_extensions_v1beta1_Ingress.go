package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_IngressList.go


// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress
// can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based
// virtual hosting etc. DEPRECATED - This group version of Ingress is deprecated by networking.k8s.io/v1beta1 Ingress. See
// the release notes for more information.
type Io_k8s_api_extensions_v1beta1_Ingress struct {
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

	// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#spec-and-status
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_IngressSpec.go
	Spec       *Io_k8s_api_extensions_v1beta1_IngressSpec       `json:"spec,omitempty"`

	// Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#spec-and-status
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_IngressStatus.go
	Status     *Io_k8s_api_extensions_v1beta1_IngressStatus     `json:"status,omitempty"`
}
