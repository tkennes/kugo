package kugo_model


// SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list
// of actions may be incomplete depending on the server's authorization mode, and any errors experienced during the
// evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason
// about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises
// confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the
// correct way to defer authorization decisions to the API server.
type Io_k8s_api_authorization_v1_SelfSubjectRulesReview struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                                `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                                `json:"kind,omitempty"`

	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta       `json:"metadata,omitempty"`

	// Spec holds information about the request being evaluated.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1_SelfSubjectRulesReviewSpec.go
	Spec       Io_k8s_api_authorization_v1_SelfSubjectRulesReviewSpec `json:"spec"`

	// Status is filled in by the server and indicates the set of actions a user can perform.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1_SubjectRulesReviewStatus.go
	Status     *Io_k8s_api_authorization_v1_SubjectRulesReviewStatus  `json:"status,omitempty"`
}
