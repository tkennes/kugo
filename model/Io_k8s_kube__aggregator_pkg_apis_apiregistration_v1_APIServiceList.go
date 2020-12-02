package kugo_model

type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceList struct {
	ApiVersion string                                                           `json:"apiVersion,omitempty"`
	Items      []Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIService `json:"items"`
	Kind       string                                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta                   `json:"metadata,omitempty"`
}

