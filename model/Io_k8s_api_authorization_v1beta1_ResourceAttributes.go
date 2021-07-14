package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReviewSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SubjectAccessReviewSpec.go


// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
type Io_k8s_api_authorization_v1beta1_ResourceAttributes struct {
	// Group is the API Group of the Resource.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Group       *string `json:"group,omitempty"`

	// Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name        *string `json:"name,omitempty"`

	// Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and
	// all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources ""
	// (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace   *string `json:"namespace,omitempty"`

	// Resource is one of the existing resource types.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Resource    *string `json:"resource,omitempty"`

	// Subresource is one of the existing resource types.  "" means none.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Subresource *string `json:"subresource,omitempty"`

	// Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Verb        *string `json:"verb,omitempty"`

	// Version is the API Version of the Resource.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Version     *string `json:"version,omitempty"`
}
