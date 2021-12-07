package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicySpec.go


// NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched by a
// NetworkPolicySpec's podSelector. The traffic must match both ports and to. This type is beta-level in 1.8
type Io_k8s_api_networking_v1_NetworkPolicyEgressRule struct {
	// List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is
	// empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at
	// least one item, then this rule allows traffic only if the traffic matches at least one port in the list.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyPort.go
	Ports []Io_k8s_api_networking_v1_NetworkPolicyPort `json:"ports,omitempty"`

	// List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a
	// logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by
	// destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic
	// matches at least one item in the to list.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyPeer.go
	To    []Io_k8s_api_networking_v1_NetworkPolicyPeer `json:"to,omitempty"`
}
