package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_APIGroup.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_APIVersions.go


// ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the
// clientCIDR that they match.
type Io_k8s_apimachinery_pkg_apis_meta_v1_ServerAddressByClientCIDR struct {
	// The CIDR with which clients can match their IP to figure out the server address that they should use.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ClientCIDR    *string `json:"clientCIDR"`

	// Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or
	// IP:port.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ServerAddress *string `json:"serverAddress"`
}
