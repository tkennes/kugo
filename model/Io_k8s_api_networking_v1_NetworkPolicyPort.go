package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyIngressRule.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyEgressRule.go


// NetworkPolicyPort describes a port to allow traffic on
type Io_k8s_api_networking_v1_NetworkPolicyPort struct {
	// The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided,
	// this matches all port names and numbers.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Port     *int    `json:"port,omitempty"`

	// The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Protocol *string `json:"protocol,omitempty"`
}
