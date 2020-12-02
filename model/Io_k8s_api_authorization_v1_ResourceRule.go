package kugo_model

type Io_k8s_api_authorization_v1_ResourceRule struct {
	ApiGroups     []string `json:"apiGroups,omitempty"`
	ResourceNames []string `json:"resourceNames,omitempty"`
	Resources     []string `json:"resources,omitempty"`
	Verbs         []string `json:"verbs"`
}

