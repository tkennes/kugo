package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachmentSpec.go


// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via
// external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type Io_k8s_api_storage_v1_VolumeAttachmentSource struct {
	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline
	// VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's
	// inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the
	// CSIMigration feature.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
	InlineVolumeSpec     *Io_k8s_api_core_v1_PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty"`

	// Name of the persistent volume to attach.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	PersistentVolumeName *string                                  `json:"persistentVolumeName,omitempty"`
}
