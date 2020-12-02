package kugo_model

type Io_k8s_api_extensions_v1beta1_IngressRule struct {
	Host string                                              `json:"host,omitempty"`
	Http *Io_k8s_api_extensions_v1beta1_HTTPIngressRuleValue `json:"http,omitempty"`
}

