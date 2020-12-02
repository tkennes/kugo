package kugo_model

type Io_k8s_api_core_v1_ResourceQuotaSpec struct {
	Hard          *interface{}                      `json:"hard,omitempty"`
	ScopeSelector *Io_k8s_api_core_v1_ScopeSelector `json:"scopeSelector,omitempty"`
	Scopes        []string                          `json:"scopes,omitempty"`
}

