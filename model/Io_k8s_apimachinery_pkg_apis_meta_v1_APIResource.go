package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_APIResourceList.go


// APIResource specifies the name of a resource and whether it is namespaced.
type Io_k8s_apimachinery_pkg_apis_meta_v1_APIResource struct {
	// categories is a list of the grouped resources this resource belongs to (e.g. 'all')
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Categories         []*string `json:"categories,omitempty"`

	// group is the preferred group of the resource.  Empty implies the group of the containing resource list. For
	// subresources, this may have a different value, for example: Scale".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Group              *string   `json:"group,omitempty"`

	// kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind               *string   `json:"kind"`

	// name is the plural name of the resource.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name               *string   `json:"name"`

	// namespaced indicates if a resource is namespaced or not.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Namespaced         *bool     `json:"namespaced"`

	// shortNames is a list of suggested short names of the resource.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ShortNames         []*string `json:"shortNames,omitempty"`

	// singularName is the singular name of the resource.  This allows clients to handle plural and singular opaquely. The
	// singularName is more correct for reporting status on a single item and both singular and plural are allowed from the
	// kubectl CLI interface.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SingularName       *string   `json:"singularName"`

	// The hash value of the storage version, the version this resource is converted to when written to the data store. Value
	// must be treated as opaque by clients. Only equality comparison on the value is valid. This is an alpha feature and may
	// change or be removed in the future. The field is populated by the apiserver only if the StorageVersionHash feature gate
	// is enabled. This field will remain optional even if it graduates.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	StorageVersionHash *string   `json:"storageVersionHash,omitempty"`

	// verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete,
	// deletecollection, and proxy)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Verbs              []*string `json:"verbs"`

	// version is the preferred version of the resource.  Empty implies the version of the containing resource list For
	// subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's
	// group)".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Version            *string   `json:"version,omitempty"`
}
