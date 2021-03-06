package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// AllowedCSIDriver represents a single inline CSI Driver that is allowed to be used.
type Io_k8s_api_policy_v1beta1_AllowedCSIDriver struct {
	// Name is the registered name of the CSI driver
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name *string `json:"name"`
}
