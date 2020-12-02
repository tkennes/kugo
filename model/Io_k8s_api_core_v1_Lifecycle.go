package kugo_model

type Io_k8s_api_core_v1_Lifecycle struct {
	PostStart *Io_k8s_api_core_v1_Handler `json:"postStart,omitempty"`
	PreStop   *Io_k8s_api_core_v1_Handler `json:"preStop,omitempty"`
}

