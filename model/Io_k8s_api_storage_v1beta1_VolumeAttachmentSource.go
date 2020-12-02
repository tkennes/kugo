package kugo_model

type Io_k8s_api_storage_v1beta1_VolumeAttachmentSource struct {
	InlineVolumeSpec     *Io_k8s_api_core_v1_PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty"`
	PersistentVolumeName string                                   `json:"persistentVolumeName,omitempty"`
}

