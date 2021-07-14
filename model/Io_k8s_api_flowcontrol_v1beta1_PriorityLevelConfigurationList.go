package kugo_model


// PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.
type Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                                     `json:"apiVersion,omitempty"`

	// `items` is a list of request-priorities.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfiguration.go
	Items      []Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfiguration `json:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                                     `json:"kind,omitempty"`

	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta              `json:"metadata,omitempty"`
}
