package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSIDriverList.go


// CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster.
// Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object
// to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.
type Io_k8s_api_storage_v1_CSIDriver struct {
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

	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the
	// same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less,
	// beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Specification of the CSI Driver.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSIDriverSpec.go
	Spec       Io_k8s_api_storage_v1_CSIDriverSpec              `json:"spec"`
}
