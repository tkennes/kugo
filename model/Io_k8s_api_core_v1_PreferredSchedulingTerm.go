package kugo_model

type Io_k8s_api_core_v1_PreferredSchedulingTerm struct {
	Preference Io_k8s_api_core_v1_NodeSelectorTerm `json:"preference"`
	Weight     int                                 `json:"weight"`
}

