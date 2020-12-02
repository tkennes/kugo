package kugo_model

type Io_k8s_api_core_v1_ContainerState struct {
	Running    *Io_k8s_api_core_v1_ContainerStateRunning    `json:"running,omitempty"`
	Terminated *Io_k8s_api_core_v1_ContainerStateTerminated `json:"terminated,omitempty"`
	Waiting    *Io_k8s_api_core_v1_ContainerStateWaiting    `json:"waiting,omitempty"`
}

