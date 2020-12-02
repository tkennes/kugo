package kugo_model

type Io_k8s_api_authorization_v1_SubjectRulesReviewStatus struct {
	EvaluationError  string                                        `json:"evaluationError,omitempty"`
	Incomplete       bool                                          `json:"incomplete"`
	NonResourceRules []Io_k8s_api_authorization_v1_NonResourceRule `json:"nonResourceRules"`
	ResourceRules    []Io_k8s_api_authorization_v1_ResourceRule    `json:"resourceRules"`
}

