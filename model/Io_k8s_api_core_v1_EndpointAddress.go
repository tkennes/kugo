package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EndpointSubset.go


// EndpointAddress is a tuple that describes single IP address.
type Io_k8s_api_core_v1_EndpointAddress struct {
	// The Hostname of this endpoint
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Hostname  *string                             `json:"hostname,omitempty"`

	// The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast
	// ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components,
	// like kube-proxy, are not IPv6 ready.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Ip        *string                             `json:"ip"`

	// Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	NodeName  *string                             `json:"nodeName,omitempty"`

	// Reference to object providing the endpoint.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	TargetRef *Io_k8s_api_core_v1_ObjectReference `json:"targetRef,omitempty"`
}
