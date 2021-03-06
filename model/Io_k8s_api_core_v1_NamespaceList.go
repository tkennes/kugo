package kugo_model


// NamespaceList is a list of Namespaces.
type Io_k8s_api_core_v1_NamespaceList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                        `json:"apiVersion,omitempty"`

	// Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-
	// with-objects/namespaces/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Namespace.go
	Items      []Io_k8s_api_core_v1_Namespace                 `json:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                        `json:"kind,omitempty"`

	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta `json:"metadata,omitempty"`
}
