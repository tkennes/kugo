package kugo_model

type Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource struct {
	FsType    string `json:"fsType,omitempty"`
	Partition int    `json:"partition,omitempty"`
	PdName    string `json:"pdName"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}

