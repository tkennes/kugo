package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SelfSubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_LocalSubjectAccessReview.go


// SubjectAccessReviewStatus
type Io_k8s_api_authorization_v1_SubjectAccessReviewStatus struct {
	// Allowed is required. True if the action would be allowed, false otherwise.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Allowed         *bool   `json:"allowed"`

	// Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false,
	// then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Denied          *bool   `json:"denied,omitempty"`

	// EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get
	// an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a
	// role, but enough roles are still present and bound to reason about the request.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	EvaluationError *string `json:"evaluationError,omitempty"`

	// Reason is optional.  It indicates why a request was allowed or denied.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason          *string `json:"reason,omitempty"`
}
