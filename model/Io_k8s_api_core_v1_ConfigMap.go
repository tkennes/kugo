package kugo_model

type Io_k8s_api_core_v1_ConfigMap struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	BinaryData *interface{}                                     `json:"binaryData,omitempty"`
	Data       *interface{}                                     `json:"data,omitempty"`
	Immutable  bool                                             `json:"immutable,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
}

