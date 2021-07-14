package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricStatus.go


// ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
type Io_k8s_api_autoscaling_v2beta1_ExternalMetricStatus struct {
	// currentAverageValue is the current value of metric averaged over autoscaled pods.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	CurrentAverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"currentAverageValue,omitempty"`

	// currentValue is the current value of the metric (as a quantity)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	CurrentValue        *Io_k8s_apimachinery_pkg_api_resource_Quantity      `json:"currentValue"`

	// metricName is the name of a metric used for autoscaling in metric system.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MetricName          *string                                             `json:"metricName"`

	// metricSelector is used to identify a specific time series within a given metric.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	MetricSelector      *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"metricSelector,omitempty"`
}
