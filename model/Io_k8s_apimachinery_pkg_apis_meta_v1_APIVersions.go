package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_APIVersions struct {
	ApiVersion                 string                                                           `json:"apiVersion,omitempty"`
	Kind                       string                                                           `json:"kind,omitempty"`
	ServerAddressByClientCIDRs []Io_k8s_apimachinery_pkg_apis_meta_v1_ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs"`
	Versions                   []string                                                         `json:"versions"`
}

