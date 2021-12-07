package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricSpec.go


// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example,
// transactions-processed-per-second). The values will be averaged together before being compared to the target value.
type Io_k8s_api_autoscaling_v2beta2_PodsMetricSource struct {
	// metric identifies the target metric by name and selector
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricIdentifier.go
	Metric Io_k8s_api_autoscaling_v2beta2_MetricIdentifier `json:"metric"`

	// target specifies the target value for the given metric
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricTarget.go
	Target Io_k8s_api_autoscaling_v2beta2_MetricTarget     `json:"target"`
}
