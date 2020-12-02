package kugo_model

type Io_k8s_api_autoscaling_v2beta1_ObjectMetricSource struct {
	AverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity             `json:"averageValue,omitempty"`
	MetricName   string                                                     `json:"metricName"`
	Selector     *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector        `json:"selector,omitempty"`
	Target       Io_k8s_api_autoscaling_v2beta1_CrossVersionObjectReference `json:"target"`
	TargetValue  Io_k8s_apimachinery_pkg_api_resource_Quantity              `json:"targetValue"`
}

