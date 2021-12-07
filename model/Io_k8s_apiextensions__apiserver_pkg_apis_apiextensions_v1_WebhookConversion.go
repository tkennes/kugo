package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceConversion.go


// WebhookConversion describes how to call a conversion webhook
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookConversion struct {
	// clientConfig is the instructions for how to call the webhook if strategy is `Webhook`.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookClientConfig.go
	ClientConfig             *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookClientConfig `json:"clientConfig,omitempty"`

	// conversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. The API server
	// will use the first version in the list which it supports. If none of the versions specified in this list are supported
	// by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed
	// versions and does not include any versions known to the API Server, calls to the webhook will fail.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ConversionReviewVersions []*string                                                                      `json:"conversionReviewVersions"`
}
