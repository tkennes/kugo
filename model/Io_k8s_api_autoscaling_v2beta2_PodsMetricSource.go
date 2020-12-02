package kugo_model

type Io_k8s_api_autoscaling_v2beta2_PodsMetricSource struct {
	Metric Io_k8s_api_autoscaling_v2beta2_MetricIdentifier `json:"metric"`
	Target Io_k8s_api_autoscaling_v2beta2_MetricTarget     `json:"target"`
}

