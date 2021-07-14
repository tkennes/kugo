package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ObjectMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_PodsMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ResourceMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ContainerResourceMetricStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ExternalMetricStatus.go


// MetricValueStatus holds the current value for a metric
type Io_k8s_api_autoscaling_v2beta2_MetricValueStatus struct {
	// currentAverageUtilization is the current value of the average of the resource metric across all relevant pods,
	// represented as a percentage of the requested value of the resource for the pods.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	AverageUtilization *int                                           `json:"averageUtilization,omitempty"`

	// averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	AverageValue       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"averageValue,omitempty"`

	// value is the current value of the metric (as a quantity).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	Value              *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"value,omitempty"`
}
