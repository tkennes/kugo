package kugo_model

type Io_k8s_api_authentication_v1beta1_TokenReview struct {
	ApiVersion string                                               `json:"apiVersion,omitempty"`
	Kind       string                                               `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta     `json:"metadata,omitempty"`
	Spec       Io_k8s_api_authentication_v1beta1_TokenReviewSpec    `json:"spec"`
	Status     *Io_k8s_api_authentication_v1beta1_TokenReviewStatus `json:"status,omitempty"`
}

