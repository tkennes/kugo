package kugo_model

type Io_k8s_api_core_v1_NFSVolumeSource struct {
	Path     string `json:"path"`
	ReadOnly bool   `json:"readOnly,omitempty"`
	Server   string `json:"server"`
}

