package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookClientConfig struct {
	CaBundle string                                                                      `json:"caBundle,omitempty"`
	Service  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_ServiceReference `json:"service,omitempty"`
	Url      string                                                                      `json:"url,omitempty"`
}

