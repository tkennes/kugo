package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyPeer.go


// IPBlock describes a particular CIDR (Ex. "192.168.1.1/24","2001:db9::/64") that is allowed to the pods matched by a
// NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
type Io_k8s_api_networking_v1_IPBlock struct {
	// CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Cidr   *string   `json:"cidr"`

	// Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or
	// "2001:db9::/64" Except values will be rejected if they are outside the CIDR range
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Except []*string `json:"except,omitempty"`
}
