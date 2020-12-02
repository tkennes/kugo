package kugo_model

type Io_k8s_api_policy_v1beta1_PodDisruptionBudget struct {
	ApiVersion string                                               `json:"apiVersion,omitempty"`
	Kind       string                                               `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta     `json:"metadata,omitempty"`
	Spec       *Io_k8s_api_policy_v1beta1_PodDisruptionBudgetSpec   `json:"spec,omitempty"`
	Status     *Io_k8s_api_policy_v1beta1_PodDisruptionBudgetStatus `json:"status,omitempty"`
}

