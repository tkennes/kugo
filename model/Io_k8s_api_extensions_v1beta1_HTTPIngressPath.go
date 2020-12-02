package kugo_model

type Io_k8s_api_extensions_v1beta1_HTTPIngressPath struct {
	Backend  Io_k8s_api_extensions_v1beta1_IngressBackend `json:"backend"`
	Path     string                                       `json:"path,omitempty"`
	PathType string                                       `json:"pathType,omitempty"`
}

