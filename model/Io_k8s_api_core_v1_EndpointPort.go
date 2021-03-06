package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EndpointSubset.go


// EndpointPort is a tuple that describes a single port.
type Io_k8s_api_core_v1_EndpointPort struct {
	// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are
	// reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-
	// standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. This is a beta field that is
	// guarded by the ServiceAppProtocol feature gate and enabled by default.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	AppProtocol *string `json:"appProtocol,omitempty"`

	// The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional
	// only if one port is defined.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name        *string `json:"name,omitempty"`

	// The port number of the endpoint.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Port        *int    `json:"port"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Protocol    *string `json:"protocol,omitempty"`
}
