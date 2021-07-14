package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricStatus.go


// ContainerResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in
// requests and limits, describing a single container in each pod in the current scale target (e.g. CPU or memory).  Such
// metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics
// using the "pods" source.
type Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricStatus struct {
	// container is the name of the container in the pods of the scaling target
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Container                 *string                                        `json:"container"`

	// currentAverageUtilization is the current value of the average of the resource metric across all relevant pods,
	// represented as a percentage of the requested value of the resource for the pods.  It will only be present if
	// `targetAverageValue` was set in the corresponding metric specification.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CurrentAverageUtilization *int                                           `json:"currentAverageUtilization,omitempty"`

	// currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value
	// (instead of as a percentage of the request), similar to the "pods" metric source type. It will always be set, regardless
	// of the corresponding metric specification.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	CurrentAverageValue       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"currentAverageValue"`

	// name is the name of the resource in question.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name                      *string                                        `json:"name"`
}
