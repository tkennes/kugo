package kugo_model

type Io_k8s_api_networking_v1_IngressRule struct {
	Host string                                         `json:"host,omitempty"`
	Http *Io_k8s_api_networking_v1_HTTPIngressRuleValue `json:"http,omitempty"`
}

