package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionSpec.go


// CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionNames struct {
	// categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API
	// discovery documents, and used by clients to support invocations like `kubectl get all`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Categories []*string `json:"categories,omitempty"`

	// kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use
	// this value as the `kind` attribute in API calls.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string   `json:"kind"`

	// listKind is the serialized kind of the list for this resource. Defaults to "`kind`List".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ListKind   *string   `json:"listKind,omitempty"`

	// plural is the plural name of the resource to serve. The custom resources are served under
	// `/apis/<group>/<version>/.../<plural>`. Must match the name of the CustomResourceDefinition (in the form
	// `<names.plural>.<group>`). Must be all lowercase.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Plural     *string   `json:"plural"`

	// shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support
	// invocations like `kubectl get <shortname>`. It must be all lowercase.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ShortNames []*string `json:"shortNames,omitempty"`

	// singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased `kind`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Singular   *string   `json:"singular,omitempty"`
}
