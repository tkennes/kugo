package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DaemonSet.go


// DaemonSetStatus represents the current status of a daemon set.
type Io_k8s_api_apps_v1_DaemonSetStatus struct {
	// Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism
	// when it needs to create the name for the newest ControllerRevision.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	CollisionCount         *int                                    `json:"collisionCount,omitempty"`

	// Represents the latest available observations of a DaemonSet's current state.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DaemonSetCondition.go
	Conditions             []Io_k8s_api_apps_v1_DaemonSetCondition `json:"conditions,omitempty"`

	// The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	CurrentNumberScheduled *int                                    `json:"currentNumberScheduled"`

	// The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More
	// info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	DesiredNumberScheduled *int                                    `json:"desiredNumberScheduled"`

	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available
	// (ready for at least spec.minReadySeconds)
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	NumberAvailable        *int                                    `json:"numberAvailable,omitempty"`

	// The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	NumberMisscheduled     *int                                    `json:"numberMisscheduled"`

	// The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	NumberReady            *int                                    `json:"numberReady"`

	// The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready
	// for at least spec.minReadySeconds)
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	NumberUnavailable      *int                                    `json:"numberUnavailable,omitempty"`

	// The most recent generation observed by the daemon set controller.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	ObservedGeneration     *int                                    `json:"observedGeneration,omitempty"`

	// The total number of nodes that are running updated daemon pod
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	UpdatedNumberScheduled *int                                    `json:"updatedNumberScheduled,omitempty"`
}
