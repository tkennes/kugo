package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricStatus.go


// ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
type Io_k8s_api_autoscaling_v2beta2_ExternalMetricStatus struct {
	// current contains the current value for the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricValueStatus.go
	Current Io_k8s_api_autoscaling_v2beta2_MetricValueStatus `json:"current"`

	// metric identifies the target metric by name and selector
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricIdentifier.go
	Metric  Io_k8s_api_autoscaling_v2beta2_MetricIdentifier  `json:"metric"`
}
