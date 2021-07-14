package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_ReplicaSet.go


// ReplicaSetStatus represents the current status of a ReplicaSet.
type Io_k8s_api_apps_v1_ReplicaSetStatus struct {
	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	AvailableReplicas    *int                                     `json:"availableReplicas,omitempty"`

	// Represents the latest available observations of a replica set's current state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_ReplicaSetCondition.go
	Conditions           []Io_k8s_api_apps_v1_ReplicaSetCondition `json:"conditions,omitempty"`

	// The number of pods that have labels matching the labels of the pod template of the replicaset.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	FullyLabeledReplicas *int                                     `json:"fullyLabeledReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration   *int                                     `json:"observedGeneration,omitempty"`

	// The number of ready replicas for this replica set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ReadyReplicas        *int                                     `json:"readyReplicas,omitempty"`

	// Replicas is the most recently oberved number of replicas. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas             *int                                     `json:"replicas"`
}
