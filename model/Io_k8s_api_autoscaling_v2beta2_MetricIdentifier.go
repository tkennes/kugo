package kugo_model

type Io_k8s_api_autoscaling_v2beta2_MetricIdentifier struct {
	Name     string                                              `json:"name"`
	Selector *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
}

