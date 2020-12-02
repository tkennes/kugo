package kugo_model

type Io_k8s_api_core_v1_ResourceFieldSelector struct {
	ContainerName string                                         `json:"containerName,omitempty"`
	Divisor       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"divisor,omitempty"`
	Resource      string                                         `json:"resource"`
}

