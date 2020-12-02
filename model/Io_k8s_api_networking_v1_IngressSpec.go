package kugo_model

type Io_k8s_api_networking_v1_IngressSpec struct {
	DefaultBackend   *Io_k8s_api_networking_v1_IngressBackend `json:"defaultBackend,omitempty"`
	IngressClassName string                                   `json:"ingressClassName,omitempty"`
	Rules            []Io_k8s_api_networking_v1_IngressRule   `json:"rules,omitempty"`
	Tls              []Io_k8s_api_networking_v1_IngressTLS    `json:"tls,omitempty"`
}

