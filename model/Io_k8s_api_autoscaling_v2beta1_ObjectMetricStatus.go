package kugo_model

type Io_k8s_api_autoscaling_v2beta1_ObjectMetricStatus struct {
	AverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity             `json:"averageValue,omitempty"`
	CurrentValue Io_k8s_apimachinery_pkg_api_resource_Quantity              `json:"currentValue"`
	MetricName   string                                                     `json:"metricName"`
	Selector     *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector        `json:"selector,omitempty"`
	Target       Io_k8s_api_autoscaling_v2beta1_CrossVersionObjectReference `json:"target"`
}

