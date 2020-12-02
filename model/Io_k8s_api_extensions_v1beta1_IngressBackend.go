package kugo_model

type Io_k8s_api_extensions_v1beta1_IngressBackend struct {
	Resource    *Io_k8s_api_core_v1_TypedLocalObjectReference `json:"resource,omitempty"`
	ServiceName string                                        `json:"serviceName,omitempty"`
	ServicePort int                                           `json:"servicePort,omitempty"`
}

