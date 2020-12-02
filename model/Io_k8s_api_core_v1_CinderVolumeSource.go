package kugo_model

type Io_k8s_api_core_v1_CinderVolumeSource struct {
	FsType    string                                   `json:"fsType,omitempty"`
	ReadOnly  bool                                     `json:"readOnly,omitempty"`
	SecretRef *Io_k8s_api_core_v1_LocalObjectReference `json:"secretRef,omitempty"`
	VolumeID  string                                   `json:"volumeID"`
}

