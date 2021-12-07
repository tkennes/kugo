package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// SELinuxStrategyOptions defines the strategy type and any options used to create the strategy.
type Io_k8s_api_policy_v1beta1_SELinuxStrategyOptions struct {
	// rule is the strategy that will dictate the allowable labels that may be set.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Rule           *string                            `json:"rule"`

	// seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-
	// container/security-context/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SELinuxOptions.go
	SeLinuxOptions *Io_k8s_api_core_v1_SELinuxOptions `json:"seLinuxOptions,omitempty"`
}
