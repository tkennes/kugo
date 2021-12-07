package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodSpec.go


// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
type Io_k8s_api_core_v1_HostAlias struct {
	// Hostnames for the above IP address.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Hostnames []*string `json:"hostnames,omitempty"`

	// IP address of the host file entry.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Ip        *string   `json:"ip,omitempty"`
}
