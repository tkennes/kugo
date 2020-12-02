package kugo_model

type Io_k8s_api_core_v1_GlusterfsPersistentVolumeSource struct {
	Endpoints          string `json:"endpoints"`
	EndpointsNamespace string `json:"endpointsNamespace,omitempty"`
	Path               string `json:"path"`
	ReadOnly           bool   `json:"readOnly,omitempty"`
}

