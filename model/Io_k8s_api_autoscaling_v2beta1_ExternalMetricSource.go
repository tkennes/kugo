package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricSpec.go


// ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of
// queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one "target" type should
// be set.
type Io_k8s_api_autoscaling_v2beta1_ExternalMetricSource struct {
	// metricName is the name of the metric in question.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	MetricName         *string                                             `json:"metricName"`

	// metricSelector is used to identify a specific time series within a given metric.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	MetricSelector     *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"metricSelector,omitempty"`

	// targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	TargetAverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"targetAverageValue,omitempty"`

	// targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	TargetValue        *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"targetValue,omitempty"`
}
