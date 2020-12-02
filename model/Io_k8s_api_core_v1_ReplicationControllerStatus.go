package kugo_model

type Io_k8s_api_core_v1_ReplicationControllerStatus struct {
	AvailableReplicas    int                                                 `json:"availableReplicas,omitempty"`
	Conditions           []Io_k8s_api_core_v1_ReplicationControllerCondition `json:"conditions,omitempty"`
	FullyLabeledReplicas int                                                 `json:"fullyLabeledReplicas,omitempty"`
	ObservedGeneration   int                                                 `json:"observedGeneration,omitempty"`
	ReadyReplicas        int                                                 `json:"readyReplicas,omitempty"`
	Replicas             int                                                 `json:"replicas"`
}

