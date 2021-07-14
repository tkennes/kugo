package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1_ValidatingWebhookConfigurationList.go


// ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object
// without changing it.
type Io_k8s_api_admissionregistration_v1_ValidatingWebhookConfiguration struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                                 `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                                 `json:"kind,omitempty"`

	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta        `json:"metadata,omitempty"`

	// Webhooks is a list of webhooks and the affected resources and operations.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1_ValidatingWebhook.go
	Webhooks   []Io_k8s_api_admissionregistration_v1_ValidatingWebhook `json:"webhooks,omitempty"`
}
