package kugo_model

type Io_k8s_api_admissionregistration_v1_MutatingWebhookConfiguration struct {
	ApiVersion string                                                `json:"apiVersion,omitempty"`
	Kind       string                                                `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta      `json:"metadata,omitempty"`
	Webhooks   []Io_k8s_api_admissionregistration_v1_MutatingWebhook `json:"webhooks,omitempty"`
}

