package model


// PersistentVolumeList is a list of PersistentVolume items.
type Io_k8s_api_core_v1_PersistentVolumeList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                        `json:"apiVersion,omitempty"`

	// List of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolume.go
	Items      []Io_k8s_api_core_v1_PersistentVolume          `json:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                        `json:"kind,omitempty"`

	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta `json:"metadata,omitempty"`
}
