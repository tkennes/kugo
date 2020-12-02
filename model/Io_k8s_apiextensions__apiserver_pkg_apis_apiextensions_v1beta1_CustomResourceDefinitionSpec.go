package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionSpec struct {
	AdditionalPrinterColumns []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceColumnDefinition  `json:"additionalPrinterColumns,omitempty"`
	Conversion               *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceConversion         `json:"conversion,omitempty"`
	Group                    string                                                                                           `json:"group"`
	Names                    Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionNames     `json:"names"`
	PreserveUnknownFields    bool                                                                                             `json:"preserveUnknownFields,omitempty"`
	Scope                    string                                                                                           `json:"scope"`
	Subresources             *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceSubresources       `json:"subresources,omitempty"`
	Validation               *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceValidation         `json:"validation,omitempty"`
	Version                  string                                                                                           `json:"version,omitempty"`
	Versions                 []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionVersion `json:"versions,omitempty"`
}

