package kugo_model

type Io_k8s_api_policy_v1beta1_Eviction struct {
	ApiVersion    string                                              `json:"apiVersion,omitempty"`
	DeleteOptions *Io_k8s_apimachinery_pkg_apis_meta_v1_DeleteOptions `json:"deleteOptions,omitempty"`
	Kind          string                                              `json:"kind,omitempty"`
	Metadata      *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta    `json:"metadata,omitempty"`
}

