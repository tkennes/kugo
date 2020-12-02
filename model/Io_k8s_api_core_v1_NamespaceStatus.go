package kugo_model

type Io_k8s_api_core_v1_NamespaceStatus struct {
	Conditions []Io_k8s_api_core_v1_NamespaceCondition `json:"conditions,omitempty"`
	Phase      string                                  `json:"phase,omitempty"`
}

