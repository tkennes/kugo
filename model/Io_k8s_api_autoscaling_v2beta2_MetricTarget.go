package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ObjectMetricSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_PodsMetricSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ResourceMetricSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ContainerResourceMetricSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_ExternalMetricSource.go


// MetricTarget defines the target value, average value, or average utilization of a specific metric
type Io_k8s_api_autoscaling_v2beta2_MetricTarget struct {
	// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a
	// percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	AverageUtilization *int                                           `json:"averageUtilization,omitempty"`

	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	AverageValue       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"averageValue,omitempty"`

	// type represents whether the metric type is Utilization, Value, or AverageValue
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string                                        `json:"type"`

	// value is the target value of the metric (as a quantity).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	Value              *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"value,omitempty"`
}
