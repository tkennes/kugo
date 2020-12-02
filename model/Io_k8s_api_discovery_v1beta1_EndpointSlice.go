package kugo_model

type Io_k8s_api_discovery_v1beta1_EndpointSlice struct {
	AddressType string                                           `json:"addressType"`
	ApiVersion  string                                           `json:"apiVersion,omitempty"`
	Endpoints   []Io_k8s_api_discovery_v1beta1_Endpoint          `json:"endpoints"`
	Kind        string                                           `json:"kind,omitempty"`
	Metadata    *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Ports       []Io_k8s_api_discovery_v1beta1_EndpointPort      `json:"ports,omitempty"`
}

