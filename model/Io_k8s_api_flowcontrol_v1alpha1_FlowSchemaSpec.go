package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_FlowSchema.go


// FlowSchemaSpec describes how the FlowSchema's specification looks like.
type Io_k8s_api_flowcontrol_v1alpha1_FlowSchemaSpec struct {
	// `distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies
	// that the distinguisher is disabled and thus will always be the empty string.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_FlowDistinguisherMethod.go
	DistinguisherMethod        *Io_k8s_api_flowcontrol_v1alpha1_FlowDistinguisherMethod            `json:"distinguisherMethod,omitempty"`

	// `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among
	// those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence
	// value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	MatchingPrecedence         *int                                                                `json:"matchingPrecedence,omitempty"`

	// `priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be
	// resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationReference.go
	PriorityLevelConfiguration Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationReference `json:"priorityLevelConfiguration"`

	// `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least
	// one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PolicyRulesWithSubjects.go
	Rules                      []Io_k8s_api_flowcontrol_v1alpha1_PolicyRulesWithSubjects           `json:"rules,omitempty"`
}
