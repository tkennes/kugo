package kugo_model

type Io_k8s_api_flowcontrol_v1alpha1_PolicyRulesWithSubjects struct {
	NonResourceRules []Io_k8s_api_flowcontrol_v1alpha1_NonResourcePolicyRule `json:"nonResourceRules,omitempty"`
	ResourceRules    []Io_k8s_api_flowcontrol_v1alpha1_ResourcePolicyRule    `json:"resourceRules,omitempty"`
	Subjects         []Io_k8s_api_flowcontrol_v1alpha1_Subject               `json:"subjects"`
}

