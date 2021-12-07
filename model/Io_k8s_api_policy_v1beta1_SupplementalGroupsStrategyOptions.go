package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
type Io_k8s_api_policy_v1beta1_SupplementalGroupsStrategyOptions struct {
	// ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then
	// supply a single range with the same start and end. Required for MustRunAs.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_IDRange.go
	Ranges []Io_k8s_api_policy_v1beta1_IDRange `json:"ranges,omitempty"`

	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Rule   *string                             `json:"rule,omitempty"`
}
