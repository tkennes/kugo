package kugo_model

type Io_k8s_api_core_v1_CephFSVolumeSource struct {
	Monitors   []string                                 `json:"monitors"`
	Path       string                                   `json:"path,omitempty"`
	ReadOnly   bool                                     `json:"readOnly,omitempty"`
	SecretFile string                                   `json:"secretFile,omitempty"`
	SecretRef  *Io_k8s_api_core_v1_LocalObjectReference `json:"secretRef,omitempty"`
	User       string                                   `json:"user,omitempty"`
}

