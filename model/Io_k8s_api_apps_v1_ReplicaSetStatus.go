package kugo_model

type Io_k8s_api_apps_v1_ReplicaSetStatus struct {
	AvailableReplicas    int                                      `json:"availableReplicas,omitempty"`
	Conditions           []Io_k8s_api_apps_v1_ReplicaSetCondition `json:"conditions,omitempty"`
	FullyLabeledReplicas int                                      `json:"fullyLabeledReplicas,omitempty"`
	ObservedGeneration   int                                      `json:"observedGeneration,omitempty"`
	ReadyReplicas        int                                      `json:"readyReplicas,omitempty"`
	Replicas             int                                      `json:"replicas"`
}

