package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SubjectRulesReviewStatus.go


// NonResourceRule holds information that describes a rule for the non-resource
type Io_k8s_api_authorization_v1_NonResourceRule struct {
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final
	// step in the path.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	NonResourceURLs []*string `json:"nonResourceURLs,omitempty"`

	// Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Verbs           []*string `json:"verbs"`
}
