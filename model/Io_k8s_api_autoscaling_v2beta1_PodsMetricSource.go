package kugo_model

type Io_k8s_api_autoscaling_v2beta1_PodsMetricSource struct {
	MetricName         string                                              `json:"metricName"`
	Selector           *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
	TargetAverageValue Io_k8s_apimachinery_pkg_api_resource_Quantity       `json:"targetAverageValue"`
}

