package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricSpec.go


// ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and
// limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together
// before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top
// of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
type Io_k8s_api_autoscaling_v2beta2_ResourceMetricSource struct {
	// name is the name of the resource in question.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name   *string                                     `json:"name"`

	// target specifies the target value for the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricTarget.go
	Target Io_k8s_api_autoscaling_v2beta2_MetricTarget `json:"target"`
}
