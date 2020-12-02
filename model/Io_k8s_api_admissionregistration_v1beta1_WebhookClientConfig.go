package kugo_model

type Io_k8s_api_admissionregistration_v1beta1_WebhookClientConfig struct {
	CaBundle string                                                     `json:"caBundle,omitempty"`
	Service  *Io_k8s_api_admissionregistration_v1beta1_ServiceReference `json:"service,omitempty"`
	Url      string                                                     `json:"url,omitempty"`
}

