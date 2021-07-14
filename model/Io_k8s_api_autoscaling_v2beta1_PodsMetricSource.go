package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricSpec.go


// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example,
// transactions-processed-per-second). The values will be averaged together before being compared to the target value.
type Io_k8s_api_autoscaling_v2beta1_PodsMetricSource struct {
	// metricName is the name of the metric in question
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MetricName         *string                                             `json:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed
	// as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will
	// be used to gather metrics.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector           *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`

	// targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	TargetAverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"targetAverageValue"`
}
