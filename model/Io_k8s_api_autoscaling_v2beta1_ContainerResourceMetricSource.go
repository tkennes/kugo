package kugo_model

type Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricSource struct {
	Container                string                                         `json:"container"`
	Name                     string                                         `json:"name"`
	TargetAverageUtilization int                                            `json:"targetAverageUtilization,omitempty"`
	TargetAverageValue       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"targetAverageValue,omitempty"`
}

