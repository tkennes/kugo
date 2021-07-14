package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodStatus.go


// ContainerStatus contains details for the current status of this container.
type Io_k8s_api_core_v1_ContainerStatus struct {
	// Container's ID in the format 'docker://<container_id>'.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ContainerID  *string                            `json:"containerID,omitempty"`

	// The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Image        *string                            `json:"image"`

	// ImageID of the container's image.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ImageID      *string                            `json:"imageID"`

	// Details about the container's last termination condition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerState.go
	LastState    *Io_k8s_api_core_v1_ContainerState `json:"lastState,omitempty"`

	// This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name         *string                            `json:"name"`

	// Specifies whether the container has passed its readiness probe.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Ready        *bool                              `json:"ready"`

	// The number of times the container has been restarted, currently based on the number of dead containers that have not yet
	// been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection.
	// This value will get capped at 5 by GC.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	RestartCount *int                               `json:"restartCount"`

	// Specifies whether the container has passed its startup probe. Initialized as false, becomes true after startupProbe is
	// considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. Is always
	// true when no startupProbe is defined.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Started      *bool                              `json:"started,omitempty"`

	// Details about the container's current condition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerState.go
	State        *Io_k8s_api_core_v1_ContainerState `json:"state,omitempty"`
}
