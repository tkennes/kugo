package kugo_model

type Io_k8s_api_core_v1_LoadBalancerIngress struct {
	Hostname string                          `json:"hostname,omitempty"`
	Ip       string                          `json:"ip,omitempty"`
	Ports    []Io_k8s_api_core_v1_PortStatus `json:"ports,omitempty"`
}

