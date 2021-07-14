package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SubjectRulesReviewStatus.go


// ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant,
// may contain duplicates, and possibly be incomplete.
type Io_k8s_api_authorization_v1beta1_ResourceRule struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action
	// requested against one of the enumerated resources in any API group will be allowed.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiGroups     []*string `json:"apiGroups,omitempty"`

	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is
	// allowed.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ResourceNames []*string `json:"resourceNames,omitempty"`

	// Resources is a list of resources this rule applies to.  "*" means all in the specified apiGroups.  "*/foo" represents
	// the subresource 'foo' for all resources in the specified apiGroups.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Resources     []*string `json:"resources,omitempty"`

	// Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  "*" means all.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Verbs         []*string `json:"verbs"`
}
