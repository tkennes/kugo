package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Container.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go


// Lifecycle describes actions that the management system should take in response to container lifecycle events. For the
// PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the
// container process fails, in which case the handler is aborted.
type Io_k8s_api_core_v1_Lifecycle struct {
	// PostStart is called immediately after a container is created. If the handler fails, the container is terminated and
	// restarted according to its restart policy. Other management of the container blocks until the hook completes. More info:
	// https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Handler.go
	PostStart *Io_k8s_api_core_v1_Handler `json:"postStart,omitempty"`

	// PreStop is called immediately before a container is terminated due to an API request or management event such as
	// liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes
	// or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins
	// before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate
	// within the Pod's termination grace period. Other management of the container blocks until the hook completes or until
	// the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-
	// hooks/#container-hooks
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Handler.go
	PreStop   *Io_k8s_api_core_v1_Handler `json:"preStop,omitempty"`
}
