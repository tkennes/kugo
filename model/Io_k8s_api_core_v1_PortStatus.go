package kugo_model

type Io_k8s_api_core_v1_PortStatus struct {
	Error    string `json:"error,omitempty"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

