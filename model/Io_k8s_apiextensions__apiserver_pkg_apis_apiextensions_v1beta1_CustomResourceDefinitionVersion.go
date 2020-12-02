package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionVersion struct {
	AdditionalPrinterColumns []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceColumnDefinition `json:"additionalPrinterColumns,omitempty"`
	Deprecated               bool                                                                                            `json:"deprecated,omitempty"`
	DeprecationWarning       string                                                                                          `json:"deprecationWarning,omitempty"`
	Name                     string                                                                                          `json:"name"`
	Schema                   *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceValidation        `json:"schema,omitempty"`
	Served                   bool                                                                                            `json:"served"`
	Storage                  bool                                                                                            `json:"storage"`
	Subresources             *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceSubresources      `json:"subresources,omitempty"`
}

