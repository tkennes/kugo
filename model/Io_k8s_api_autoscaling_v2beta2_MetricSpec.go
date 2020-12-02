package kugo_model

type Io_k8s_api_autoscaling_v2beta2_MetricSpec struct {
	ContainerResource *Io_k8s_api_autoscaling_v2beta2_ContainerResourceMetricSource `json:"containerResource,omitempty"`
	External          *Io_k8s_api_autoscaling_v2beta2_ExternalMetricSource          `json:"external,omitempty"`
	Object            *Io_k8s_api_autoscaling_v2beta2_ObjectMetricSource            `json:"object,omitempty"`
	Pods              *Io_k8s_api_autoscaling_v2beta2_PodsMetricSource              `json:"pods,omitempty"`
	Resource          *Io_k8s_api_autoscaling_v2beta2_ResourceMetricSource          `json:"resource,omitempty"`
	Type              string                                                        `json:"type"`
}

