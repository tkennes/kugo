package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerStatus.go


// MetricStatus describes the last-read state of a single metric.
type Io_k8s_api_autoscaling_v2beta1_MetricStatus struct {
	// container resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes
	// describing a single container in each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to
	// Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods"
	// source.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricStatus.go
	ContainerResource *Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricStatus `json:"containerResource,omitempty"`

	// external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on
	// information coming from components running outside of cluster (for example length of queue in cloud messaging service,
	// or QPS from loadbalancer running outside of cluster).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_ExternalMetricStatus.go
	External          *Io_k8s_api_autoscaling_v2beta1_ExternalMetricStatus          `json:"external,omitempty"`

	// object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_ObjectMetricStatus.go
	Object            *Io_k8s_api_autoscaling_v2beta1_ObjectMetricStatus            `json:"object,omitempty"`

	// pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-
	// second).  The values will be averaged together before being compared to the target value.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_PodsMetricStatus.go
	Pods              *Io_k8s_api_autoscaling_v2beta1_PodsMetricStatus              `json:"pods,omitempty"`

	// resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing
	// each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special
	// scaling options on top of those available to normal per-pod metrics using the "pods" source.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_ResourceMetricStatus.go
	Resource          *Io_k8s_api_autoscaling_v2beta1_ResourceMetricStatus          `json:"resource,omitempty"`

	// type is the type of metric source.  It will be one of "ContainerResource", "External", "Object", "Pods" or "Resource",
	// each corresponds to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate
	// HPAContainerMetrics is enabled
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type              *string                                                       `json:"type"`
}
