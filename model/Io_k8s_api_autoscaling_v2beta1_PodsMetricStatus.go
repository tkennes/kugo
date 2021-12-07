package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricStatus.go


// PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example,
// transactions-processed-per-second).
type Io_k8s_api_autoscaling_v2beta1_PodsMetricStatus struct {
	// currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	CurrentAverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"currentAverageValue"`

	// metricName is the name of the metric in question
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	MetricName          *string                                             `json:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the
	// PodsMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When
	// unset, just the metricName will be used to gather metrics.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector            *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
}
