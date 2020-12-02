package kugo_model

type Io_k8s_api_rbac_v1_PolicyRule struct {
	ApiGroups       []string `json:"apiGroups,omitempty"`
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty"`
	Resources       []string `json:"resources,omitempty"`
	Verbs           []string `json:"verbs"`
}

