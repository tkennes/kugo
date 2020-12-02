package kugo_model

type Io_k8s_api_core_v1_SecretProjection struct {
	Items    []Io_k8s_api_core_v1_KeyToPath `json:"items,omitempty"`
	Name     string                         `json:"name,omitempty"`
	Optional bool                           `json:"optional,omitempty"`
}

