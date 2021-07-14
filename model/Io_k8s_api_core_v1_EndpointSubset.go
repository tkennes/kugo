package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Endpoints.go


// EndpointSubset is a group of addresses with a common set of ports. The expanded set of endpoints is the Cartesian
// product of Addresses x Ports. For example, given:   {     Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
// Ports:     [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]   } The resulting set of endpoints can be viewed
// as:     a: [ 10.10.1.1:8675, 10.10.2.2:8675 ],     b: [ 10.10.1.1:309, 10.10.2.2:309 ]
type Io_k8s_api_core_v1_EndpointSubset struct {
	// IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load
	// balancers and clients to utilize.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EndpointAddress.go
	Addresses         []Io_k8s_api_core_v1_EndpointAddress `json:"addresses,omitempty"`

	// IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished
	// starting, have recently failed a readiness check, or have recently failed a liveness check.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EndpointAddress.go
	NotReadyAddresses []Io_k8s_api_core_v1_EndpointAddress `json:"notReadyAddresses,omitempty"`

	// Port numbers available on the related IP addresses.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EndpointPort.go
	Ports             []Io_k8s_api_core_v1_EndpointPort    `json:"ports,omitempty"`
}
