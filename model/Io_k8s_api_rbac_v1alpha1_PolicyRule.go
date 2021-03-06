package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1alpha1_Role.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1alpha1_ClusterRole.go


// PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies
// to or which namespace the rule applies to.
type Io_k8s_api_rbac_v1alpha1_PolicyRule struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action
	// requested against one of the enumerated resources in any API group will be allowed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiGroups       []*string `json:"apiGroups,omitempty"`

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final
	// step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced
	// from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL
	// paths (such as "/api"),  but not both.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	NonResourceURLs []*string `json:"nonResourceURLs,omitempty"`

	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is
	// allowed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ResourceNames   []*string `json:"resourceNames,omitempty"`

	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Resources       []*string `json:"resources,omitempty"`

	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll
	// represents all kinds.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Verbs           []*string `json:"verbs"`
}
