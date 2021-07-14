package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricStatus.go


// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second
// on an Ingress object).
type Io_k8s_api_autoscaling_v2beta2_ObjectMetricStatus struct {
	// current contains the current value for the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricValueStatus.go
	Current         Io_k8s_api_autoscaling_v2beta2_MetricValueStatus           `json:"current"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_CrossVersionObjectReference.go
	DescribedObject Io_k8s_api_autoscaling_v2beta2_CrossVersionObjectReference `json:"describedObject"`

	// metric identifies the target metric by name and selector
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricIdentifier.go
	Metric          Io_k8s_api_autoscaling_v2beta2_MetricIdentifier            `json:"metric"`
}
