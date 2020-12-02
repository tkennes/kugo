package kugo_model

type Io_k8s_api_core_v1_FlexPersistentVolumeSource struct {
	Driver    string                              `json:"driver"`
	FsType    string                              `json:"fsType,omitempty"`
	Options   *interface{}                        `json:"options,omitempty"`
	ReadOnly  bool                                `json:"readOnly,omitempty"`
	SecretRef *Io_k8s_api_core_v1_SecretReference `json:"secretRef,omitempty"`
}

