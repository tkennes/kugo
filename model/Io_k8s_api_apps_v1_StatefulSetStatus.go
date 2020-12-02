package kugo_model

type Io_k8s_api_apps_v1_StatefulSetStatus struct {
	CollisionCount     int                                       `json:"collisionCount,omitempty"`
	Conditions         []Io_k8s_api_apps_v1_StatefulSetCondition `json:"conditions,omitempty"`
	CurrentReplicas    int                                       `json:"currentReplicas,omitempty"`
	CurrentRevision    string                                    `json:"currentRevision,omitempty"`
	ObservedGeneration int                                       `json:"observedGeneration,omitempty"`
	ReadyReplicas      int                                       `json:"readyReplicas,omitempty"`
	Replicas           int                                       `json:"replicas"`
	UpdateRevision     string                                    `json:"updateRevision,omitempty"`
	UpdatedReplicas    int                                       `json:"updatedReplicas,omitempty"`
}

