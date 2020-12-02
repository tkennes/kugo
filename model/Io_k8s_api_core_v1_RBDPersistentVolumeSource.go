package kugo_model

type Io_k8s_api_core_v1_RBDPersistentVolumeSource struct {
	FsType    string                              `json:"fsType,omitempty"`
	Image     string                              `json:"image"`
	Keyring   string                              `json:"keyring,omitempty"`
	Monitors  []string                            `json:"monitors"`
	Pool      string                              `json:"pool,omitempty"`
	ReadOnly  bool                                `json:"readOnly,omitempty"`
	SecretRef *Io_k8s_api_core_v1_SecretReference `json:"secretRef,omitempty"`
	User      string                              `json:"user,omitempty"`
}

