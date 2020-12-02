package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_APIResourceList struct {
	ApiVersion   string                                             `json:"apiVersion,omitempty"`
	GroupVersion string                                             `json:"groupVersion"`
	Kind         string                                             `json:"kind,omitempty"`
	Resources    []Io_k8s_apimachinery_pkg_apis_meta_v1_APIResource `json:"resources"`
}

