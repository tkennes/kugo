package kugo_model

type Io_k8s_api_core_v1_Endpoints struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Subsets    []Io_k8s_api_core_v1_EndpointSubset              `json:"subsets,omitempty"`
}

