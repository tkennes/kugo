package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_APIGroup.go


// GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.
type Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery struct {
	// groupVersion specifies the API group and version in the form "group/version"
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	GroupVersion *string `json:"groupVersion"`

	// version specifies the version in the form of "version". This is to save the clients the trouble of splitting the
	// GroupVersion.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Version      *string `json:"version"`
}
