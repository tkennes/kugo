package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_StorageClassList.go


// StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
// StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
type Io_k8s_api_storage_v1beta1_StorageClass struct {
	// AllowVolumeExpansion shows whether the storage class allow volume expand
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	AllowVolumeExpansion *bool                                            `json:"allowVolumeExpansion,omitempty"`

	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported
	// topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only
	// honored by servers that enable the VolumeScheduling feature.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_TopologySelectorTerm.go
	AllowedTopologies    []Io_k8s_api_core_v1_TopologySelectorTerm        `json:"allowedTopologies,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion           *string                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind                 *string                                          `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata             *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro",
	// "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	MountOptions         []*string                                        `json:"mountOptions,omitempty"`

	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Parameters           *interface{}                                     `json:"parameters,omitempty"`

	// Provisioner indicates the type of the provisioner.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Provisioner          *string                                          `json:"provisioner"`

	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ReclaimPolicy        *string                                          `json:"reclaimPolicy,omitempty"`

	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset,
	// VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	VolumeBindingMode    *string                                          `json:"volumeBindingMode,omitempty"`
}
