package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionStatus struct {
	AcceptedNames  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionNames      `json:"acceptedNames,omitempty"`
	Conditions     []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionCondition `json:"conditions,omitempty"`
	StoredVersions []string                                                                                      `json:"storedVersions,omitempty"`
}

