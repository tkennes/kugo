package kugo_model

type Io_k8s_api_flowcontrol_v1alpha1_ResourcePolicyRule struct {
	ApiGroups    []string `json:"apiGroups"`
	ClusterScope bool     `json:"clusterScope,omitempty"`
	Namespaces   []string `json:"namespaces,omitempty"`
	Resources    []string `json:"resources"`
	Verbs        []string `json:"verbs"`
}

