package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_Deployment.go


// DeploymentStatus is the most recently observed status of the Deployment.
type Io_k8s_api_apps_v1_DeploymentStatus struct {
	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	AvailableReplicas   *int                                     `json:"availableReplicas,omitempty"`

	// Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance
	// mechanism when it needs to create the name for the newest ReplicaSet.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CollisionCount      *int                                     `json:"collisionCount,omitempty"`

	// Represents the latest available observations of a deployment's current state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_DeploymentCondition.go
	Conditions          []Io_k8s_api_apps_v1_DeploymentCondition `json:"conditions,omitempty"`

	// The generation observed by the deployment controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration  *int                                     `json:"observedGeneration,omitempty"`

	// Total number of ready pods targeted by this deployment.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ReadyReplicas       *int                                     `json:"readyReplicas,omitempty"`

	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas            *int                                     `json:"replicas,omitempty"`

	// Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required
	// for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or
	// pods that still have not been created.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	UnavailableReplicas *int                                     `json:"unavailableReplicas,omitempty"`

	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	UpdatedReplicas     *int                                     `json:"updatedReplicas,omitempty"`
}
