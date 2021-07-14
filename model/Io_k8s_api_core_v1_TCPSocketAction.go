package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Handler.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Probe.go


// TCPSocketAction describes an action based on opening a socket
type Io_k8s_api_core_v1_TCPSocketAction struct {
	// Optional: Host name to connect to, defaults to the pod IP.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Host *string `json:"host,omitempty"`

	// Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an
	// IANA_SVC_NAME.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Port *int    `json:"port"`
}
