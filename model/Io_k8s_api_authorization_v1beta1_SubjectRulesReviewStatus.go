package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SelfSubjectRulesReview.go


// SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of
// authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are
// additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is
// incomplete.
type Io_k8s_api_authorization_v1beta1_SubjectRulesReviewStatus struct {
	// EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an
	// authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	EvaluationError  *string                                            `json:"evaluationError,omitempty"`

	// Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an
	// authorizer, such as an external authorizer, doesn't support rules evaluation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Incomplete       *bool                                              `json:"incomplete"`

	// NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't
	// significant, may contain duplicates, and possibly be incomplete.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_NonResourceRule.go
	NonResourceRules []Io_k8s_api_authorization_v1beta1_NonResourceRule `json:"nonResourceRules"`

	// ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't
	// significant, may contain duplicates, and possibly be incomplete.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_ResourceRule.go
	ResourceRules    []Io_k8s_api_authorization_v1beta1_ResourceRule    `json:"resourceRules"`
}
