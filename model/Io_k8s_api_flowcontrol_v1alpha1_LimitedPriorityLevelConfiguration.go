package kugo_model

type Io_k8s_api_flowcontrol_v1alpha1_LimitedPriorityLevelConfiguration struct {
	AssuredConcurrencyShares int                                            `json:"assuredConcurrencyShares,omitempty"`
	LimitResponse            *Io_k8s_api_flowcontrol_v1alpha1_LimitResponse `json:"limitResponse,omitempty"`
}

