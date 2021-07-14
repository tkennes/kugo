package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_ClusterRoleList.go


// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or
// ClusterRoleBinding.
type Io_k8s_api_rbac_v1_ClusterRole struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is
	// set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_AggregationRule.go
	AggregationRule *Io_k8s_api_rbac_v1_AggregationRule              `json:"aggregationRule,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion      *string                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind            *string                                          `json:"kind,omitempty"`

	// Standard object's metadata.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata        *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Rules holds all the PolicyRules for this ClusterRole
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_PolicyRule.go
	Rules           []Io_k8s_api_rbac_v1_PolicyRule                  `json:"rules,omitempty"`
}
