package kugo_model

type Io_k8s_api_core_v1_ServiceAccount struct {
	ApiVersion                   string                                           `json:"apiVersion,omitempty"`
	AutomountServiceAccountToken bool                                             `json:"automountServiceAccountToken,omitempty"`
	ImagePullSecrets             []Io_k8s_api_core_v1_LocalObjectReference        `json:"imagePullSecrets,omitempty"`
	Kind                         string                                           `json:"kind,omitempty"`
	Metadata                     *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Secrets                      []Io_k8s_api_core_v1_ObjectReference             `json:"secrets,omitempty"`
}

