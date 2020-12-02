package kugo_model

type Io_k8s_api_core_v1_WeightedPodAffinityTerm struct {
	PodAffinityTerm Io_k8s_api_core_v1_PodAffinityTerm `json:"podAffinityTerm"`
	Weight          int                                `json:"weight"`
}

