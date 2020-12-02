package kugo_model

type Io_k8s_api_networking_v1_IngressBackend struct {
	Resource *Io_k8s_api_core_v1_TypedLocalObjectReference   `json:"resource,omitempty"`
	Service  *Io_k8s_api_networking_v1_IngressServiceBackend `json:"service,omitempty"`
}

