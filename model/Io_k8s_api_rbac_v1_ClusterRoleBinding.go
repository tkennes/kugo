package kugo_model

type Io_k8s_api_rbac_v1_ClusterRoleBinding struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	RoleRef    Io_k8s_api_rbac_v1_RoleRef                       `json:"roleRef"`
	Subjects   []Io_k8s_api_rbac_v1_Subject                     `json:"subjects,omitempty"`
}

