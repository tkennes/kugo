package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_EndpointSliceList.go


// EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple
// EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
type Io_k8s_api_discovery_v1beta1_EndpointSlice struct {
	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same
	// type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an
	// IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	AddressType *string                                          `json:"addressType"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion  *string                                          `json:"apiVersion,omitempty"`

	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_Endpoint.go
	Endpoints   []Io_k8s_api_discovery_v1beta1_Endpoint          `json:"endpoints"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind        *string                                          `json:"kind,omitempty"`

	// Standard object's metadata.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata    *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name.
	// When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it
	// indicates "all ports". Each slice may include a maximum of 100 ports.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_EndpointPort.go
	Ports       []Io_k8s_api_discovery_v1beta1_EndpointPort      `json:"ports,omitempty"`
}
