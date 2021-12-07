package kugo_model


// LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a
// namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.
type Io_k8s_api_authorization_v1_LocalSubjectAccessReview struct {
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

	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the
	// request against.  If empty, it is defaulted.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1_SubjectAccessReviewSpec.go
	Spec       Io_k8s_api_authorization_v1_SubjectAccessReviewSpec    `json:"spec"`

	// Status is filled in by the server and indicates whether the request is allowed or not
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1_SubjectAccessReviewStatus.go
	Status     *Io_k8s_api_authorization_v1_SubjectAccessReviewStatus `json:"status,omitempty"`
}
