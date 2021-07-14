package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricSpec.go


// ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests
// and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together
// before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top
// of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
type Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricSource struct {
	// container is the name of the container in the pods of the scaling target
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Container                *string                                        `json:"container"`

	// name is the name of the resource in question.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name                     *string                                        `json:"name"`

	// targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented
	// as a percentage of the requested value of the resource for the pods.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	TargetAverageUtilization *int                                           `json:"targetAverageUtilization,omitempty"`

	// targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value
	// (instead of as a percentage of the request), similar to the "pods" metric source type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	TargetAverageValue       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"targetAverageValue,omitempty"`
}
