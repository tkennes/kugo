package kugo_model

type Io_k8s_api_autoscaling_v2beta1_MetricSpec struct {
	ContainerResource *Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricSource `json:"containerResource,omitempty"`
	External          *Io_k8s_api_autoscaling_v2beta1_ExternalMetricSource          `json:"external,omitempty"`
	Object            *Io_k8s_api_autoscaling_v2beta1_ObjectMetricSource            `json:"object,omitempty"`
	Pods              *Io_k8s_api_autoscaling_v2beta1_PodsMetricSource              `json:"pods,omitempty"`
	Resource          *Io_k8s_api_autoscaling_v2beta1_ResourceMetricSource          `json:"resource,omitempty"`
	Type              string                                                        `json:"type"`
}

