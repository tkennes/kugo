package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionList.go


//   Storage version of a specific resource.
type Io_k8s_api_apiserverinternal_v1alpha1_StorageVersion struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                                    `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                                    `json:"kind,omitempty"`

	// The name is <group>.<resource>.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta           `json:"metadata,omitempty"`

	// Spec is an empty spec. It is here to comply with Kubernetes API style.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionSpec.go
	Spec       *Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionSpec  `json:"spec"`

	// API server instances report the version they can decode and the version they encode objects to when persisting objects
	// in the backend.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionStatus.go
	Status     Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionStatus `json:"status"`
}
