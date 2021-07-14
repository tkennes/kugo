package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricStatus.go


// ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and
// limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes,
// and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
type Io_k8s_api_autoscaling_v2beta2_ResourceMetricStatus struct {
	// current contains the current value for the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricValueStatus.go
	Current Io_k8s_api_autoscaling_v2beta2_MetricValueStatus `json:"current"`

	// Name is the name of the resource in question.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name    *string                                          `json:"name"`
}
