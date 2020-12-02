package kugo_model

type Io_k8s_api_core_v1_NodeSpec struct {
	ConfigSource  *Io_k8s_api_core_v1_NodeConfigSource `json:"configSource,omitempty"`
	ExternalID    string                               `json:"externalID,omitempty"`
	PodCIDR       string                               `json:"podCIDR,omitempty"`
	PodCIDRs      []string                             `json:"podCIDRs,omitempty"`
	ProviderID    string                               `json:"providerID,omitempty"`
	Taints        []Io_k8s_api_core_v1_Taint           `json:"taints,omitempty"`
	Unschedulable bool                                 `json:"unschedulable,omitempty"`
}

