package kugo_model

type Io_k8s_api_apps_v1_ControllerRevision struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Data       *Io_k8s_apimachinery_pkg_runtime_RawExtension    `json:"data,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Revision   int                                              `json:"revision"`
}

