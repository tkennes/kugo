package kugo_model

type Io_k8s_api_core_v1_PodDNSConfig struct {
	Nameservers []string                                `json:"nameservers,omitempty"`
	Options     []Io_k8s_api_core_v1_PodDNSConfigOption `json:"options,omitempty"`
	Searches    []string                                `json:"searches,omitempty"`
}

