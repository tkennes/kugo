package kugo_model

type Io_k8s_api_policy_v1beta1_PodDisruptionBudgetSpec struct {
	MaxUnavailable int                                                 `json:"maxUnavailable,omitempty"`
	MinAvailable   int                                                 `json:"minAvailable,omitempty"`
	Selector       *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
}

