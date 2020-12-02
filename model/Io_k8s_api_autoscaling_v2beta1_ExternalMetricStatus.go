package kugo_model

type Io_k8s_api_autoscaling_v2beta1_ExternalMetricStatus struct {
	CurrentAverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"currentAverageValue,omitempty"`
	CurrentValue        Io_k8s_apimachinery_pkg_api_resource_Quantity       `json:"currentValue"`
	MetricName          string                                              `json:"metricName"`
	MetricSelector      *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"metricSelector,omitempty"`
}

