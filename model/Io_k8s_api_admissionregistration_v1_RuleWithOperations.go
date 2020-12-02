package kugo_model

type Io_k8s_api_admissionregistration_v1_RuleWithOperations struct {
	ApiGroups   []string `json:"apiGroups,omitempty"`
	ApiVersions []string `json:"apiVersions,omitempty"`
	Operations  []string `json:"operations,omitempty"`
	Resources   []string `json:"resources,omitempty"`
	Scope       string   `json:"scope,omitempty"`
}

