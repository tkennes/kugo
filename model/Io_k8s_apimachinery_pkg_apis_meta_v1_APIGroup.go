package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_APIGroupList.go


// APIGroup contains the name, the supported versions, and the preferred version of a group.
type Io_k8s_apimachinery_pkg_apis_meta_v1_APIGroup struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion                 *string                                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind                       *string                                                          `json:"kind,omitempty"`

	// name is the name of the group.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name                       *string                                                          `json:"name"`

	// preferredVersion is the version preferred by the API server, which probably is the storage version.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery.go
	PreferredVersion           *Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery   `json:"preferredVersion,omitempty"`

	// a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most
	// network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case
	// of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks
	// that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the
	// server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that
	// order) to get the client IP.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ServerAddressByClientCIDR.go
	ServerAddressByClientCIDRs []Io_k8s_apimachinery_pkg_apis_meta_v1_ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs,omitempty"`

	// versions are the versions supported in this group.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery.go
	Versions                   []Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery  `json:"versions"`
}
