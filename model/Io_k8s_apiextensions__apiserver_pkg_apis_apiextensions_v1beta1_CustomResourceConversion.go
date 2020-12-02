package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceConversion struct {
	ConversionReviewVersions []string                                                                            `json:"conversionReviewVersions,omitempty"`
	Strategy                 string                                                                              `json:"strategy"`
	WebhookClientConfig      *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_WebhookClientConfig `json:"webhookClientConfig,omitempty"`
}

