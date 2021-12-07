package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ReplicationController.go


// ReplicationControllerSpec is the specification of a replication controller.
type Io_k8s_api_core_v1_ReplicationControllerSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to
	// be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	MinReadySeconds *int                                `json:"minReadySeconds,omitempty"`

	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified.
	// Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-
	// replicationcontroller
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Replicas        *int                                `json:"replicas,omitempty"`

	// Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the
	// labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication
	// controller, if empty defaulted to labels on Pod template. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Selector        *interface{}                        `json:"selector,omitempty"`

	// Template is the object that describes the pod that will be created if insufficient replicas are detected. This takes
	// precedence over a TemplateRef. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodTemplateSpec.go
	Template        *Io_k8s_api_core_v1_PodTemplateSpec `json:"template,omitempty"`
}
