package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricSpec.go


// ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an
// Ingress object).
type Io_k8s_api_autoscaling_v2beta2_ObjectMetricSource struct {
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_CrossVersionObjectReference.go
	DescribedObject Io_k8s_api_autoscaling_v2beta2_CrossVersionObjectReference `json:"describedObject"`

	// metric identifies the target metric by name and selector
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricIdentifier.go
	Metric          Io_k8s_api_autoscaling_v2beta2_MetricIdentifier            `json:"metric"`

	// target specifies the target value for the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricTarget.go
	Target          Io_k8s_api_autoscaling_v2beta2_MetricTarget                `json:"target"`
}
