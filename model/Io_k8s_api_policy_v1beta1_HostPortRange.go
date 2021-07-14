package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// HostPortRange defines a range of host ports that will be enabled by a policy for pods to use.  It requires both the
// start and end to be defined.
type Io_k8s_api_policy_v1beta1_HostPortRange struct {
	// max is the end of the range, inclusive.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Max *int `json:"max"`

	// min is the start of the range, inclusive.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Min *int `json:"min"`
}
