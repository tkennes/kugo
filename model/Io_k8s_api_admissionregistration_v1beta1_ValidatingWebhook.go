package kugo_model

type Io_k8s_api_admissionregistration_v1beta1_ValidatingWebhook struct {
	AdmissionReviewVersions []string                                                      `json:"admissionReviewVersions,omitempty"`
	ClientConfig            Io_k8s_api_admissionregistration_v1beta1_WebhookClientConfig  `json:"clientConfig"`
	FailurePolicy           string                                                        `json:"failurePolicy,omitempty"`
	MatchPolicy             string                                                        `json:"matchPolicy,omitempty"`
	Name                    string                                                        `json:"name"`
	NamespaceSelector       *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector           `json:"namespaceSelector,omitempty"`
	ObjectSelector          *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector           `json:"objectSelector,omitempty"`
	Rules                   []Io_k8s_api_admissionregistration_v1beta1_RuleWithOperations `json:"rules,omitempty"`
	SideEffects             string                                                        `json:"sideEffects,omitempty"`
	TimeoutSeconds          int                                                           `json:"timeoutSeconds,omitempty"`
}

