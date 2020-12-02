package kugo_model

type Io_k8s_api_core_v1_SecretVolumeSource struct {
	DefaultMode int                            `json:"defaultMode,omitempty"`
	Items       []Io_k8s_api_core_v1_KeyToPath `json:"items,omitempty"`
	Optional    bool                           `json:"optional,omitempty"`
	SecretName  string                         `json:"secretName,omitempty"`
}

