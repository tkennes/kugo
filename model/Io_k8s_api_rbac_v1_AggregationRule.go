package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1_ClusterRole.go


// AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
type Io_k8s_api_rbac_v1_AggregationRule struct {
	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of
	// the selectors match, then the ClusterRole's permissions will be added
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	ClusterRoleSelectors []Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"clusterRoleSelectors,omitempty"`
}
