package kugo_model

type Io_k8s_api_core_v1_PersistentVolumeClaimTemplate struct {
	Metadata *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Spec     Io_k8s_api_core_v1_PersistentVolumeClaimSpec     `json:"spec"`
}

