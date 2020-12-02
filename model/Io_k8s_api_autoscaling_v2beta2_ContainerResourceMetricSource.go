package kugo_model

type Io_k8s_api_autoscaling_v2beta2_ContainerResourceMetricSource struct {
	Container string                                      `json:"container"`
	Name      string                                      `json:"name"`
	Target    Io_k8s_api_autoscaling_v2beta2_MetricTarget `json:"target"`
}

