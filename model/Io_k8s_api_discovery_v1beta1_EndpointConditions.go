package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_Endpoint.go


// EndpointConditions represents the current condition of an endpoint.
type Io_k8s_api_discovery_v1beta1_EndpointConditions struct {
	// ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the
	// endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
	// For compatibility reasons, ready should never be "true" for terminating endpoints.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Ready       *bool `json:"ready,omitempty"`

	// serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition
	// should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition.
	// This field can be enabled with the EndpointSliceTerminatingCondition feature gate.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Serving     *bool `json:"serving,omitempty"`

	// terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should
	// interpret this unknown state to mean that the endpoint is not terminating. This field can be enabled with the
	// EndpointSliceTerminatingCondition feature gate.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Terminating *bool `json:"terminating,omitempty"`
}
