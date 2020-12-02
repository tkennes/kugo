package kugo_model

type Io_k8s_api_core_v1_ContainerStatus struct {
	ContainerID  string                             `json:"containerID,omitempty"`
	Image        string                             `json:"image"`
	ImageID      string                             `json:"imageID"`
	LastState    *Io_k8s_api_core_v1_ContainerState `json:"lastState,omitempty"`
	Name         string                             `json:"name"`
	Ready        bool                               `json:"ready"`
	RestartCount int                                `json:"restartCount"`
	Started      bool                               `json:"started,omitempty"`
	State        *Io_k8s_api_core_v1_ContainerState `json:"state,omitempty"`
}

