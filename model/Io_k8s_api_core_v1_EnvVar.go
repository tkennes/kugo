package kugo_model

type Io_k8s_api_core_v1_EnvVar struct {
	Name      string                           `json:"name"`
	Value     string                           `json:"value,omitempty"`
	ValueFrom *Io_k8s_api_core_v1_EnvVarSource `json:"valueFrom,omitempty"`
}

