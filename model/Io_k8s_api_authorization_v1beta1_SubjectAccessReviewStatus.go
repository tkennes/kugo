package kugo_model

type Io_k8s_api_authorization_v1beta1_SubjectAccessReviewStatus struct {
	Allowed         bool   `json:"allowed"`
	Denied          bool   `json:"denied,omitempty"`
	EvaluationError string `json:"evaluationError,omitempty"`
	Reason          string `json:"reason,omitempty"`
}

