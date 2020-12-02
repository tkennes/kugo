package kugo_model

type Io_k8s_api_autoscaling_v2beta1_MetricStatus struct {
	ContainerResource *Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricStatus `json:"containerResource,omitempty"`
	External          *Io_k8s_api_autoscaling_v2beta1_ExternalMetricStatus          `json:"external,omitempty"`
	Object            *Io_k8s_api_autoscaling_v2beta1_ObjectMetricStatus            `json:"object,omitempty"`
	Pods              *Io_k8s_api_autoscaling_v2beta1_PodsMetricStatus              `json:"pods,omitempty"`
	Resource          *Io_k8s_api_autoscaling_v2beta1_ResourceMetricStatus          `json:"resource,omitempty"`
	Type              string                                                        `json:"type"`
}

