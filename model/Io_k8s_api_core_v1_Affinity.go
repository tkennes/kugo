package kugo_model

type Io_k8s_api_core_v1_Affinity struct {
	NodeAffinity    *Io_k8s_api_core_v1_NodeAffinity    `json:"nodeAffinity,omitempty"`
	PodAffinity     *Io_k8s_api_core_v1_PodAffinity     `json:"podAffinity,omitempty"`
	PodAntiAffinity *Io_k8s_api_core_v1_PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

