package kugo_model

type Io_k8s_api_apps_v1_DeploymentStatus struct {
	AvailableReplicas   int                                      `json:"availableReplicas,omitempty"`
	CollisionCount      int                                      `json:"collisionCount,omitempty"`
	Conditions          []Io_k8s_api_apps_v1_DeploymentCondition `json:"conditions,omitempty"`
	ObservedGeneration  int                                      `json:"observedGeneration,omitempty"`
	ReadyReplicas       int                                      `json:"readyReplicas,omitempty"`
	Replicas            int                                      `json:"replicas,omitempty"`
	UnavailableReplicas int                                      `json:"unavailableReplicas,omitempty"`
	UpdatedReplicas     int                                      `json:"updatedReplicas,omitempty"`
}

