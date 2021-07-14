package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodDisruptionBudget.go


// PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual
// state of a system.
type Io_k8s_api_policy_v1beta1_PodDisruptionBudgetStatus struct {
	// current number of healthy pods
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CurrentHealthy     *int         `json:"currentHealthy"`

	// minimum desired number of healthy pods
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DesiredHealthy     *int         `json:"desiredHealthy"`

	// DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource
	// handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time
	// when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been
	// marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the
	// API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from
	// the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be
	// empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	DisruptedPods      *interface{} `json:"disruptedPods,omitempty"`

	// Number of pod disruptions that are currently allowed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DisruptionsAllowed *int         `json:"disruptionsAllowed"`

	// total number of pods counted by this disruption budget
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ExpectedPods       *int         `json:"expectedPods"`

	// Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid
	// only if observedGeneration equals to PDB's object generation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration *int         `json:"observedGeneration,omitempty"`
}
