package kugo_model

type Io_k8s_api_autoscaling_v2beta2_MetricValueStatus struct {
	AverageUtilization int                                            `json:"averageUtilization,omitempty"`
	AverageValue       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"averageValue,omitempty"`
	Value              *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"value,omitempty"`
}

