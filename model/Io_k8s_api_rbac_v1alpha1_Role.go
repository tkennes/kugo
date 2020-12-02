package kugo_model

type Io_k8s_api_rbac_v1alpha1_Role struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Rules      []Io_k8s_api_rbac_v1alpha1_PolicyRule            `json:"rules,omitempty"`
}

