package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinition struct {
	ApiVersion string                                                                                         `json:"apiVersion,omitempty"`
	Kind       string                                                                                         `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta                                               `json:"metadata,omitempty"`
	Spec       Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionSpec    `json:"spec"`
	Status     *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionStatus `json:"status,omitempty"`
}

