package kugo_model

type Io_k8s_api_autoscaling_v2beta1_ContainerResourceMetricStatus struct {
	Container                 string                                        `json:"container"`
	CurrentAverageUtilization int                                           `json:"currentAverageUtilization,omitempty"`
	CurrentAverageValue       Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"currentAverageValue"`
	Name                      string                                        `json:"name"`
}

