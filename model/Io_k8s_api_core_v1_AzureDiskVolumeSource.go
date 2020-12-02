package kugo_model

type Io_k8s_api_core_v1_AzureDiskVolumeSource struct {
	CachingMode string `json:"cachingMode,omitempty"`
	DiskName    string `json:"diskName"`
	DiskURI     string `json:"diskURI"`
	FsType      string `json:"fsType,omitempty"`
	Kind        string `json:"kind,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
}

