package kugo_model

type Io_k8s_api_core_v1_FCVolumeSource struct {
	FsType     string   `json:"fsType,omitempty"`
	Lun        int      `json:"lun,omitempty"`
	ReadOnly   bool     `json:"readOnly,omitempty"`
	TargetWWNs []string `json:"targetWWNs,omitempty"`
	Wwids      []string `json:"wwids,omitempty"`
}

