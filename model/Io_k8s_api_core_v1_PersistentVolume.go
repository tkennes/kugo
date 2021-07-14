package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeList.go


// PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info:
// https://kubernetes.io/docs/concepts/storage/persistent-volumes
type Io_k8s_api_core_v1_PersistentVolume struct {
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

	// Spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
	Spec       *Io_k8s_api_core_v1_PersistentVolumeSpec         `json:"spec,omitempty"`

	// Status represents the current information/status for the persistent volume. Populated by the system. Read-only. More
	// info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeStatus.go
	Status     *Io_k8s_api_core_v1_PersistentVolumeStatus       `json:"status,omitempty"`
}
