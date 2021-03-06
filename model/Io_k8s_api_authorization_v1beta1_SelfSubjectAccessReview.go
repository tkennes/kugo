package kugo_model


// SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means
// "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an
// action
type Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReview struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                                      `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                                      `json:"kind,omitempty"`

	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta             `json:"metadata,omitempty"`

	// Spec holds information about the request being evaluated.  user and groups must be empty
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReviewSpec.go
	Spec       Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReviewSpec `json:"spec"`

	// Status is filled in by the server and indicates whether the request is allowed or not
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1beta1_SubjectAccessReviewStatus.go
	Status     *Io_k8s_api_authorization_v1beta1_SubjectAccessReviewStatus  `json:"status,omitempty"`
}
