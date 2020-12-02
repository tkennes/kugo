package kugo_model

type Io_k8s_api_core_v1_NodeAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []Io_k8s_api_core_v1_PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	RequiredDuringSchedulingIgnoredDuringExecution  *Io_k8s_api_core_v1_NodeSelector             `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

