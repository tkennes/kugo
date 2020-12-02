package kugo_model

type Io_k8s_api_extensions_v1beta1_IngressSpec struct {
	Backend          *Io_k8s_api_extensions_v1beta1_IngressBackend `json:"backend,omitempty"`
	IngressClassName string                                        `json:"ingressClassName,omitempty"`
	Rules            []Io_k8s_api_extensions_v1beta1_IngressRule   `json:"rules,omitempty"`
	Tls              []Io_k8s_api_extensions_v1beta1_IngressTLS    `json:"tls,omitempty"`
}

