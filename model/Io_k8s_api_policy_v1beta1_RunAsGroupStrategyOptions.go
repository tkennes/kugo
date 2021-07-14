package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// RunAsGroupStrategyOptions defines the strategy type and any options used to create the strategy.
type Io_k8s_api_policy_v1beta1_RunAsGroupStrategyOptions struct {
	// ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single
	// range with the same start and end. Required for MustRunAs.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_IDRange.go
	Ranges []Io_k8s_api_policy_v1beta1_IDRange `json:"ranges,omitempty"`

	// rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Rule   *string                             `json:"rule"`
}
