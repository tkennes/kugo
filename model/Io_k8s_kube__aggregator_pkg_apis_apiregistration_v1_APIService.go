package kugo_model

type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIService struct {
	ApiVersion string                                                                `json:"apiVersion,omitempty"`
	Kind       string                                                                `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta                      `json:"metadata,omitempty"`
	Spec       *Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceSpec   `json:"spec,omitempty"`
	Status     *Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceStatus `json:"status,omitempty"`
}

