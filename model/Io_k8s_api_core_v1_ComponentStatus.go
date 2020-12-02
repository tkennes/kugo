package kugo_model

type Io_k8s_api_core_v1_ComponentStatus struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Conditions []Io_k8s_api_core_v1_ComponentCondition          `json:"conditions,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
}

