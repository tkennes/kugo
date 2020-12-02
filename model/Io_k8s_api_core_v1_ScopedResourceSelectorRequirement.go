package kugo_model

type Io_k8s_api_core_v1_ScopedResourceSelectorRequirement struct {
	Operator  string   `json:"operator"`
	ScopeName string   `json:"scopeName"`
	Values    []string `json:"values,omitempty"`
}

