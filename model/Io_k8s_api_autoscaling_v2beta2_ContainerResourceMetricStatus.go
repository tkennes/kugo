package kugo_model

type Io_k8s_api_autoscaling_v2beta2_ContainerResourceMetricStatus struct {
	Container string                                           `json:"container"`
	Current   Io_k8s_api_autoscaling_v2beta2_MetricValueStatus `json:"current"`
	Name      string                                           `json:"name"`
}

