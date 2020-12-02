package kugo_model

type Io_k8s_api_core_v1_ComponentStatusList struct {
	ApiVersion string                                         `json:"apiVersion,omitempty"`
	Items      []Io_k8s_api_core_v1_ComponentStatus           `json:"items"`
	Kind       string                                         `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta `json:"metadata,omitempty"`
}

