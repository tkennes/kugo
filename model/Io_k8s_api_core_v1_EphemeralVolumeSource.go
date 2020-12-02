package kugo_model

type Io_k8s_api_core_v1_EphemeralVolumeSource struct {
	ReadOnly            bool                                              `json:"readOnly,omitempty"`
	VolumeClaimTemplate *Io_k8s_api_core_v1_PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"`
}

