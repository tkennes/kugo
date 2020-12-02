package kugo_model

type Io_k8s_api_core_v1_NodeConfigStatus struct {
	Active        *Io_k8s_api_core_v1_NodeConfigSource `json:"active,omitempty"`
	Assigned      *Io_k8s_api_core_v1_NodeConfigSource `json:"assigned,omitempty"`
	Error         string                               `json:"error,omitempty"`
	LastKnownGood *Io_k8s_api_core_v1_NodeConfigSource `json:"lastKnownGood,omitempty"`
}

