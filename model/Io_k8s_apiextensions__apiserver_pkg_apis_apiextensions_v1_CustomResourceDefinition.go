package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionList.go


// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format
// <.spec.name>.<.spec.group>.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinition struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                                                                   `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                                                                   `json:"kind,omitempty"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta                                          `json:"metadata,omitempty"`

	// spec describes how the user wants the resources to appear
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionSpec.go
	Spec       Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionSpec    `json:"spec"`

	// status indicates the actual state of the CustomResourceDefinition
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionStatus.go
	Status     *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionStatus `json:"status,omitempty"`
}
