package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Container.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go


// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive
// traffic.
type Io_k8s_api_core_v1_Probe struct {
	// One and only one of the following should be specified. Exec specifies the action to take.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ExecAction.go
	Exec                *Io_k8s_api_core_v1_ExecAction      `json:"exec,omitempty"`

	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value
	// is 1.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	FailureThreshold    *int                                `json:"failureThreshold,omitempty"`

	// Number of seconds after the container has started before liveness probes are initiated. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	InitialDelaySeconds *int                                `json:"initialDelaySeconds,omitempty"`

	// How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	PeriodSeconds       *int                                `json:"periodSeconds,omitempty"`

	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1
	// for liveness and startup. Minimum value is 1.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	SuccessThreshold    *int                                `json:"successThreshold,omitempty"`

	// TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_TCPSocketAction.go
	TcpSocket           *Io_k8s_api_core_v1_TCPSocketAction `json:"tcpSocket,omitempty"`

	// Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	TimeoutSeconds      *int                                `json:"timeoutSeconds,omitempty"`
}
