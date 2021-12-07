package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_admissionregistration_v1beta1_MutatingWebhookConfigurationList.go


// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change
// the object. Deprecated in v1.16, planned for removal in v1.19. Use admissionregistration.k8s.io/v1
// MutatingWebhookConfiguration instead.
type Io_k8s_api_admissionregistration_v1beta1_MutatingWebhookConfiguration struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                                    `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                                    `json:"kind,omitempty"`

	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta           `json:"metadata,omitempty"`

	// Webhooks is a list of webhooks and the affected resources and operations.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_admissionregistration_v1beta1_MutatingWebhook.go
	Webhooks   []Io_k8s_api_admissionregistration_v1beta1_MutatingWebhook `json:"webhooks,omitempty"`
}
