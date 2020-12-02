package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_APIGroup struct {
	ApiVersion                 string                                                           `json:"apiVersion,omitempty"`
	Kind                       string                                                           `json:"kind,omitempty"`
	Name                       string                                                           `json:"name"`
	PreferredVersion           *Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery   `json:"preferredVersion,omitempty"`
	ServerAddressByClientCIDRs []Io_k8s_apimachinery_pkg_apis_meta_v1_ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs,omitempty"`
	Versions                   []Io_k8s_apimachinery_pkg_apis_meta_v1_GroupVersionForDiscovery  `json:"versions"`
}

