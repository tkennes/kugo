package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_EndpointSlice.go


// Endpoint represents a single logical "backend" implementing a service.
type Io_k8s_api_discovery_v1beta1_Endpoint struct {
	// addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice
	// addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This
	// must contain at least one address but no more than 100.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Addresses  []*string                                        `json:"addresses"`

	// conditions contains information about the current status of the endpoint.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_EndpointConditions.go
	Conditions *Io_k8s_api_discovery_v1beta1_EndpointConditions `json:"conditions,omitempty"`

	// hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other
	// (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A
	// values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Hostname   *string                                          `json:"hostname,omitempty"`

	// nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.
	// This field can be enabled with the EndpointSliceNodeName feature gate.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	NodeName   *string                                          `json:"nodeName,omitempty"`

	// targetRef is a reference to a Kubernetes object that represents this endpoint.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	TargetRef  *Io_k8s_api_core_v1_ObjectReference              `json:"targetRef,omitempty"`

	// topology contains arbitrary topology information associated with the endpoint. These key/value pairs must conform with
	// the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a
	// maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: *
	// kubernetes.io/hostname: the value indicates the hostname of the node   where the endpoint is located. This should match
	// the corresponding   node label. * topology.kubernetes.io/zone: the value indicates the zone where the   endpoint is
	// located. This should match the corresponding node label. * topology.kubernetes.io/region: the value indicates the region
	// where the   endpoint is located. This should match the corresponding node label. This field is deprecated and will be
	// removed in future api versions.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Topology   *interface{}                                     `json:"topology,omitempty"`
}
