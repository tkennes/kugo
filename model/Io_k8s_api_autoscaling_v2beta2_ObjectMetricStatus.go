package kugo_model

type Io_k8s_api_autoscaling_v2beta2_ObjectMetricStatus struct {
	Current         Io_k8s_api_autoscaling_v2beta2_MetricValueStatus           `json:"current"`
	DescribedObject Io_k8s_api_autoscaling_v2beta2_CrossVersionObjectReference `json:"describedObject"`
	Metric          Io_k8s_api_autoscaling_v2beta2_MetricIdentifier            `json:"metric"`
}

