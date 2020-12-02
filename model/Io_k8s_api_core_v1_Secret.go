package kugo_model

type Io_k8s_api_core_v1_Secret struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Data       *interface{}                                     `json:"data,omitempty"`
	Immutable  bool                                             `json:"immutable,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	StringData *interface{}                                     `json:"stringData,omitempty"`
	Type       string                                           `json:"type,omitempty"`
}

