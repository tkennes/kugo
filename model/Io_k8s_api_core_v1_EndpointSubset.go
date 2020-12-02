package kugo_model

type Io_k8s_api_core_v1_EndpointSubset struct {
	Addresses         []Io_k8s_api_core_v1_EndpointAddress `json:"addresses,omitempty"`
	NotReadyAddresses []Io_k8s_api_core_v1_EndpointAddress `json:"notReadyAddresses,omitempty"`
	Ports             []Io_k8s_api_core_v1_EndpointPort    `json:"ports,omitempty"`
}

