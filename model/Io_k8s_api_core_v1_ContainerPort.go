package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Container.go


// ContainerPort represents a network port in a single container.
type Io_k8s_api_core_v1_ContainerPort struct {
	// Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ContainerPort *int    `json:"containerPort"`

	// What host IP to bind the external port to.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	HostIP        *string `json:"hostIP,omitempty"`

	// Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is
	// specified, this must match ContainerPort. Most containers do not need this.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	HostPort      *int    `json:"hostPort,omitempty"`

	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name.
	// Name for the port that can be referred to by services.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name          *string `json:"name,omitempty"`

	// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Protocol      *string `json:"protocol,omitempty"`
}
