package kugo_model

type Io_k8s_api_core_v1_NodeSelectorTerm struct {
	MatchExpressions []Io_k8s_api_core_v1_NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchFields      []Io_k8s_api_core_v1_NodeSelectorRequirement `json:"matchFields,omitempty"`
}

