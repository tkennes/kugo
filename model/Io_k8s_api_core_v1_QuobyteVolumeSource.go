package kugo_model

type Io_k8s_api_core_v1_QuobyteVolumeSource struct {
	Group    string `json:"group,omitempty"`
	ReadOnly bool   `json:"readOnly,omitempty"`
	Registry string `json:"registry"`
	Tenant   string `json:"tenant,omitempty"`
	User     string `json:"user,omitempty"`
	Volume   string `json:"volume"`
}

