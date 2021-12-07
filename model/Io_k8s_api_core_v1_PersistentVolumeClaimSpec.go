package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimTemplate.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaim.go


// PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific
// attributes
type Io_k8s_api_core_v1_PersistentVolumeClaimSpec struct {
	// AccessModes contains the desired access modes the volume should have. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	AccessModes      []*string                                           `json:"accessModes,omitempty"`

	// This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) *
	// An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order
	// to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If
	// the provisioner or an external controller can support the specified data source, it will create a new volume based on
	// the contents of the specified data source.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_TypedLocalObjectReference.go
	DataSource       *Io_k8s_api_core_v1_TypedLocalObjectReference       `json:"dataSource,omitempty"`

	// Resources represents the minimum resources the volume should have. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ResourceRequirements.go
	Resources        *Io_k8s_api_core_v1_ResourceRequirements            `json:"resources,omitempty"`

	// A label query over volumes to consider for binding.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector         *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`

	// Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-
	// volumes#class-1
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	StorageClassName *string                                             `json:"storageClassName,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in
	// claim spec.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	VolumeMode       *string                                             `json:"volumeMode,omitempty"`

	// VolumeName is the binding reference to the PersistentVolume backing this claim.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	VolumeName       *string                                             `json:"volumeName,omitempty"`
}
