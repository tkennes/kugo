package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionSpec.go


// CustomResourceConversion describes how to convert different versions of a CR.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceConversion struct {
	// strategy specifies how custom resources are converted between versions. Allowed values are: - `None`: The converter only
	// change the apiVersion and would not touch any other field in the custom resource. - `Webhook`: API Server will call to
	// an external webhook to do the conversion. Additional information   is needed for this option. This requires
	// spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Strategy *string                                                                      `json:"strategy"`

	// webhook describes how to call the conversion webhook. Required when `strategy` is set to `Webhook`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookConversion.go
	Webhook  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookConversion `json:"webhook,omitempty"`
}
