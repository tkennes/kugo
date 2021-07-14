package kugo_model


// Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A
// request to cause such an eviction is created by POSTing to .../pods/<pod name>/evictions.
type Io_k8s_api_policy_v1beta1_Eviction struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion    *string                                             `json:"apiVersion,omitempty"`

	// DeleteOptions may be provided
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_DeleteOptions.go
	DeleteOptions *Io_k8s_apimachinery_pkg_apis_meta_v1_DeleteOptions `json:"deleteOptions,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind          *string                                             `json:"kind,omitempty"`

	// ObjectMeta describes the pod that is being evicted.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata      *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta    `json:"metadata,omitempty"`
}
