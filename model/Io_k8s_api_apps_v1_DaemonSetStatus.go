package kugo_model

type Io_k8s_api_apps_v1_DaemonSetStatus struct {
	CollisionCount         int                                     `json:"collisionCount,omitempty"`
	Conditions             []Io_k8s_api_apps_v1_DaemonSetCondition `json:"conditions,omitempty"`
	CurrentNumberScheduled int                                     `json:"currentNumberScheduled"`
	DesiredNumberScheduled int                                     `json:"desiredNumberScheduled"`
	NumberAvailable        int                                     `json:"numberAvailable,omitempty"`
	NumberMisscheduled     int                                     `json:"numberMisscheduled"`
	NumberReady            int                                     `json:"numberReady"`
	NumberUnavailable      int                                     `json:"numberUnavailable,omitempty"`
	ObservedGeneration     int                                     `json:"observedGeneration,omitempty"`
	UpdatedNumberScheduled int                                     `json:"updatedNumberScheduled,omitempty"`
}

