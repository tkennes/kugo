package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// AllowedFlexVolume represents a single Flexvolume that is allowed to be used.
type Io_k8s_api_policy_v1beta1_AllowedFlexVolume struct {
	// driver is the name of the Flexvolume driver.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Driver *string `json:"driver"`
}
