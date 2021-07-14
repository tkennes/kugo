package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricSpec.go


// ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of
// queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
type Io_k8s_api_autoscaling_v2beta2_ExternalMetricSource struct {
	// metric identifies the target metric by name and selector
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricIdentifier.go
	Metric Io_k8s_api_autoscaling_v2beta2_MetricIdentifier `json:"metric"`

	// target specifies the target value for the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricTarget.go
	Target Io_k8s_api_autoscaling_v2beta2_MetricTarget     `json:"target"`
}
