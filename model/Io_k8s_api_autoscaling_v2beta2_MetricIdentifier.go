package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ObjectMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ObjectMetricSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_PodsMetricSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_PodsMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ExternalMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ExternalMetricSource.go


// MetricIdentifier defines the name and optionally selector for a metric
type Io_k8s_api_autoscaling_v2beta2_MetricIdentifier struct {
	// name is the name of the given metric
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name     *string                                             `json:"name"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed
	// as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will
	// be used to gather metrics.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
}
