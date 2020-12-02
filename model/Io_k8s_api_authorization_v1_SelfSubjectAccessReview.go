package kugo_model

type Io_k8s_api_authorization_v1_SelfSubjectAccessReview struct {
	ApiVersion string                                                  `json:"apiVersion,omitempty"`
	Kind       string                                                  `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta        `json:"metadata,omitempty"`
	Spec       Io_k8s_api_authorization_v1_SelfSubjectAccessReviewSpec `json:"spec"`
	Status     *Io_k8s_api_authorization_v1_SubjectAccessReviewStatus  `json:"status,omitempty"`
}

