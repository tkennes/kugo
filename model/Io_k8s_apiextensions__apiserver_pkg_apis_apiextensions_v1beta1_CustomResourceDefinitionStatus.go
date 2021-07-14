package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinition.go


// CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionStatus struct {
	// acceptedNames are the names that are actually being used to serve discovery. They may be different than the names in
	// spec.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionNames.go
	AcceptedNames  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionNames      `json:"acceptedNames,omitempty"`

	// conditions indicate state for particular aspects of a CustomResourceDefinition
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionCondition.go
	Conditions     []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionCondition `json:"conditions,omitempty"`

	// storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a
	// migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to
	// another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list.
	// Versions may not be removed from `spec.versions` while they exist in this list.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	StoredVersions []*string                                                                                          `json:"storedVersions,omitempty"`
}
