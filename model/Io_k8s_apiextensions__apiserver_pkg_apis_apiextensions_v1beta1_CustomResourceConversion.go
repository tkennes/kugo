package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionSpec.go


// CustomResourceConversion describes how to convert different versions of a CR.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceConversion struct {
	// conversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. The API server
	// will use the first version in the list which it supports. If none of the versions specified in this list are supported
	// by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed
	// versions and does not include any versions known to the API Server, calls to the webhook will fail. Defaults to
	// `["v1beta1"]`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ConversionReviewVersions []*string                                                                           `json:"conversionReviewVersions,omitempty"`

	// strategy specifies how custom resources are converted between versions. Allowed values are: - `None`: The converter only
	// change the apiVersion and would not touch any other field in the custom resource. - `Webhook`: API Server will call to
	// an external webhook to do the conversion. Additional information   is needed for this option. This requires
	// spec.preserveUnknownFields to be false, and spec.conversion.webhookClientConfig to be set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Strategy                 *string                                                                             `json:"strategy"`

	// webhookClientConfig is the instructions for how to call the webhook if strategy is `Webhook`. Required when `strategy`
	// is set to `Webhook`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_WebhookClientConfig.go
	WebhookClientConfig      *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_WebhookClientConfig `json:"webhookClientConfig,omitempty"`
}
