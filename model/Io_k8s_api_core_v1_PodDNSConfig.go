package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSpec.go


// PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
type Io_k8s_api_core_v1_PodDNSConfig struct {
	// A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy.
	// Duplicated nameservers will be removed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Nameservers []*string                               `json:"nameservers,omitempty"`

	// A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries
	// will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodDNSConfigOption.go
	Options     []Io_k8s_api_core_v1_PodDNSConfigOption `json:"options,omitempty"`

	// A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from
	// DNSPolicy. Duplicated search paths will be removed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Searches    []*string                               `json:"searches,omitempty"`
}
