package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReview.go


// SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and
// NonResourceAuthorizationAttributes must be set
type Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReviewSpec struct {
	// NonResourceAttributes describes information for a non-resource access request
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_NonResourceAttributes.go
	NonResourceAttributes *Io_k8s_api_authorization_v1beta1_NonResourceAttributes `json:"nonResourceAttributes,omitempty"`

	// ResourceAuthorizationAttributes describes information for a resource access request
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_ResourceAttributes.go
	ResourceAttributes    *Io_k8s_api_authorization_v1beta1_ResourceAttributes    `json:"resourceAttributes,omitempty"`
}
