package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ContainerState.go


// ContainerStateWaiting is a waiting state of a container.
type Io_k8s_api_core_v1_ContainerStateWaiting struct {
	// Message regarding why the container is not yet running.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message *string `json:"message,omitempty"`

	// (brief) reason the container is not yet running.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Reason  *string `json:"reason,omitempty"`
}
