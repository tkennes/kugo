package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeStatus.go


// NodeAddress contains information for the node's address.
type Io_k8s_api_core_v1_NodeAddress struct {
	// The node address.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Address *string `json:"address"`

	// Node address type, one of Hostname, ExternalIP or InternalIP.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type    *string `json:"type"`
}
