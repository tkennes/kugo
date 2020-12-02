package kugo_model

type Io_k8s_api_core_v1_CSIVolumeSource struct {
	Driver               string                                   `json:"driver"`
	FsType               string                                   `json:"fsType,omitempty"`
	NodePublishSecretRef *Io_k8s_api_core_v1_LocalObjectReference `json:"nodePublishSecretRef,omitempty"`
	ReadOnly             bool                                     `json:"readOnly,omitempty"`
	VolumeAttributes     *interface{}                             `json:"volumeAttributes,omitempty"`
}

