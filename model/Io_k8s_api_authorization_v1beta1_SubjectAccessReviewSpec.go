package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1beta1_SubjectAccessReview.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1beta1_LocalSubjectAccessReview.go


// SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and
// NonResourceAuthorizationAttributes must be set
type Io_k8s_api_authorization_v1beta1_SubjectAccessReviewSpec struct {
	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it
	// needs a reflection here.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Extra                 *interface{}                                            `json:"extra,omitempty"`

	// Groups is the groups you're testing for.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Group                 []*string                                               `json:"group,omitempty"`

	// NonResourceAttributes describes information for a non-resource access request
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1beta1_NonResourceAttributes.go
	NonResourceAttributes *Io_k8s_api_authorization_v1beta1_NonResourceAttributes `json:"nonResourceAttributes,omitempty"`

	// ResourceAuthorizationAttributes describes information for a resource access request
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_authorization_v1beta1_ResourceAttributes.go
	ResourceAttributes    *Io_k8s_api_authorization_v1beta1_ResourceAttributes    `json:"resourceAttributes,omitempty"`

	// UID information about the requesting user.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Uid                   *string                                                 `json:"uid,omitempty"`

	// User is the user you're testing for. If you specify "User" but not "Group", then is it interpreted as "What if User were
	// not a member of any groups
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	User                  *string                                                 `json:"user,omitempty"`
}
