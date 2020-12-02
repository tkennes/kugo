package kugo_model

type Io_k8s_api_autoscaling_v2beta2_MetricStatus struct {
	ContainerResource *Io_k8s_api_autoscaling_v2beta2_ContainerResourceMetricStatus `json:"containerResource,omitempty"`
	External          *Io_k8s_api_autoscaling_v2beta2_ExternalMetricStatus          `json:"external,omitempty"`
	Object            *Io_k8s_api_autoscaling_v2beta2_ObjectMetricStatus            `json:"object,omitempty"`
	Pods              *Io_k8s_api_autoscaling_v2beta2_PodsMetricStatus              `json:"pods,omitempty"`
	Resource          *Io_k8s_api_autoscaling_v2beta2_ResourceMetricStatus          `json:"resource,omitempty"`
	Type              string                                                        `json:"type"`
}

