package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_RunAsGroupStrategyOptions.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_SupplementalGroupsStrategyOptions.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_FSGroupStrategyOptions.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_RunAsUserStrategyOptions.go


// IDRange provides a min/max of an allowed range of IDs.
type Io_k8s_api_policy_v1beta1_IDRange struct {
	// max is the end of the range, inclusive.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Max *int `json:"max"`

	// min is the start of the range, inclusive.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Min *int `json:"min"`
}
