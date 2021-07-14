package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_StatefulSet.go


// StatefulSetStatus represents the current state of a StatefulSet.
type Io_k8s_api_apps_v1_StatefulSetStatus struct {
	// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a
	// collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CollisionCount     *int                                      `json:"collisionCount,omitempty"`

	// Represents the latest available observations of a statefulset's current state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_StatefulSetCondition.go
	Conditions         []Io_k8s_api_apps_v1_StatefulSetCondition `json:"conditions,omitempty"`

	// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by
	// currentRevision.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CurrentReplicas    *int                                      `json:"currentReplicas,omitempty"`

	// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence
	// [0,currentReplicas).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	CurrentRevision    *string                                   `json:"currentRevision,omitempty"`

	// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's
	// generation, which is updated on mutation by the API Server.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration *int                                      `json:"observedGeneration,omitempty"`

	// readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ReadyReplicas      *int                                      `json:"readyReplicas,omitempty"`

	// replicas is the number of Pods created by the StatefulSet controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas           *int                                      `json:"replicas"`

	// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-
	// updatedReplicas,replicas)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	UpdateRevision     *string                                   `json:"updateRevision,omitempty"`

	// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by
	// updateRevision.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	UpdatedReplicas    *int                                      `json:"updatedReplicas,omitempty"`
}
