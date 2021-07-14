package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerStatus.go


// ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is
// specified, the default one is ContainerStateWaiting.
type Io_k8s_api_core_v1_ContainerState struct {
	// Details about a running container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerStateRunning.go
	Running    *Io_k8s_api_core_v1_ContainerStateRunning    `json:"running,omitempty"`

	// Details about a terminated container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerStateTerminated.go
	Terminated *Io_k8s_api_core_v1_ContainerStateTerminated `json:"terminated,omitempty"`

	// Details about a waiting container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerStateWaiting.go
	Waiting    *Io_k8s_api_core_v1_ContainerStateWaiting    `json:"waiting,omitempty"`
}
