package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ReplicationController.go


// ReplicationControllerStatus represents the current status of a replication controller.
type Io_k8s_api_core_v1_ReplicationControllerStatus struct {
	// The number of available replicas (ready for at least minReadySeconds) for this replication controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	AvailableReplicas    *int                                                `json:"availableReplicas,omitempty"`

	// Represents the latest available observations of a replication controller's current state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ReplicationControllerCondition.go
	Conditions           []Io_k8s_api_core_v1_ReplicationControllerCondition `json:"conditions,omitempty"`

	// The number of pods that have labels matching the labels of the pod template of the replication controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	FullyLabeledReplicas *int                                                `json:"fullyLabeledReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed replication controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration   *int                                                `json:"observedGeneration,omitempty"`

	// The number of ready replicas for this replication controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ReadyReplicas        *int                                                `json:"readyReplicas,omitempty"`

	// Replicas is the most recently oberved number of replicas. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas             *int                                                `json:"replicas"`
}
