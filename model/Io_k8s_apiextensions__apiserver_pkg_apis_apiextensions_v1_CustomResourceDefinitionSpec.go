package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionSpec struct {
	Conversion            *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceConversion         `json:"conversion,omitempty"`
	Group                 string                                                                                      `json:"group"`
	Names                 Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionNames     `json:"names"`
	PreserveUnknownFields bool                                                                                        `json:"preserveUnknownFields,omitempty"`
	Scope                 string                                                                                      `json:"scope"`
	Versions              []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionVersion `json:"versions"`
}

