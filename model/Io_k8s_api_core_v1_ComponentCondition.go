package kugo_model

type Io_k8s_api_core_v1_ComponentCondition struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

