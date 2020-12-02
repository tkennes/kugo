package kugo_model

type Io_k8s_api_core_v1_ProjectedVolumeSource struct {
	DefaultMode int                                   `json:"defaultMode,omitempty"`
	Sources     []Io_k8s_api_core_v1_VolumeProjection `json:"sources,omitempty"`
}

