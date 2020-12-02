package kugo_model

type Io_k8s_api_core_v1_StorageOSPersistentVolumeSource struct {
	FsType          string                              `json:"fsType,omitempty"`
	ReadOnly        bool                                `json:"readOnly,omitempty"`
	SecretRef       *Io_k8s_api_core_v1_ObjectReference `json:"secretRef,omitempty"`
	VolumeName      string                              `json:"volumeName,omitempty"`
	VolumeNamespace string                              `json:"volumeNamespace,omitempty"`
}

