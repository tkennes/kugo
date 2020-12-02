package kugo_model

type Io_k8s_api_core_v1_GlusterfsVolumeSource struct {
	Endpoints string `json:"endpoints"`
	Path      string `json:"path"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}

