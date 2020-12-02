package kugo_model

type Io_k8s_api_flowcontrol_v1beta1_FlowSchemaSpec struct {
	DistinguisherMethod        *Io_k8s_api_flowcontrol_v1beta1_FlowDistinguisherMethod            `json:"distinguisherMethod,omitempty"`
	MatchingPrecedence         int                                                                `json:"matchingPrecedence,omitempty"`
	PriorityLevelConfiguration Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationReference `json:"priorityLevelConfiguration"`
	Rules                      []Io_k8s_api_flowcontrol_v1beta1_PolicyRulesWithSubjects           `json:"rules,omitempty"`
}

