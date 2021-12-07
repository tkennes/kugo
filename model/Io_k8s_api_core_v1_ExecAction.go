package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Handler.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Probe.go


// ExecAction describes a "run in container" action.
type Io_k8s_api_core_v1_ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the
	// container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions
	// ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as
	// live/healthy and non-zero is unhealthy.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Command []*string `json:"command,omitempty"`
}
