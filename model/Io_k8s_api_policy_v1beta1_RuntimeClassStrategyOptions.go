package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// RuntimeClassStrategyOptions define the strategy that will dictate the allowable RuntimeClasses for a pod.
type Io_k8s_api_policy_v1beta1_RuntimeClassStrategyOptions struct {
	// allowedRuntimeClassNames is an allowlist of RuntimeClass names that may be specified on a pod. A value of "*" means that
	// any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName
	// field to be unset.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	AllowedRuntimeClassNames []*string `json:"allowedRuntimeClassNames"`

	// defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the
	// allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	DefaultRuntimeClassName  *string   `json:"defaultRuntimeClassName,omitempty"`
}
