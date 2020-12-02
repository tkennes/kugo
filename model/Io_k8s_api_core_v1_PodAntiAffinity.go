package kugo_model

type Io_k8s_api_core_v1_PodAntiAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []Io_k8s_api_core_v1_WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	RequiredDuringSchedulingIgnoredDuringExecution  []Io_k8s_api_core_v1_PodAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

