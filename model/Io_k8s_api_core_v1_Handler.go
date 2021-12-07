package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Lifecycle.go


// Handler defines a specific action that should be taken
type Io_k8s_api_core_v1_Handler struct {
	// One and only one of the following should be specified. Exec specifies the action to take.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ExecAction.go
	Exec      *Io_k8s_api_core_v1_ExecAction      `json:"exec,omitempty"`

	// TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_TCPSocketAction.go
	TcpSocket *Io_k8s_api_core_v1_TCPSocketAction `json:"tcpSocket,omitempty"`
}
