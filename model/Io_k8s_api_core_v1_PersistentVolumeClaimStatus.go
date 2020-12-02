package kugo_model

type Io_k8s_api_core_v1_PersistentVolumeClaimStatus struct {
	AccessModes []string                                            `json:"accessModes,omitempty"`
	Capacity    *interface{}                                        `json:"capacity,omitempty"`
	Conditions  []Io_k8s_api_core_v1_PersistentVolumeClaimCondition `json:"conditions,omitempty"`
	Phase       string                                              `json:"phase,omitempty"`
}

