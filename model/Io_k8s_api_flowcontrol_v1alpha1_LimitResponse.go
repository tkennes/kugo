package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_LimitedPriorityLevelConfiguration.go


// LimitResponse defines how to handle requests that can not be executed right now.
type Io_k8s_api_flowcontrol_v1alpha1_LimitResponse struct {
	// `queuing` holds the configuration parameters for queuing. This field may be non-empty only if `type` is `"Queue"`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_QueuingConfiguration.go
	Queuing *Io_k8s_api_flowcontrol_v1alpha1_QueuingConfiguration `json:"queuing,omitempty"`

	// `type` is "Queue" or "Reject". "Queue" means that requests that can not be executed upon arrival are held in a queue
	// until they can be executed or a queuing limit is reached. "Reject" means that requests that can not be executed upon
	// arrival are rejected. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type    *string                                               `json:"type"`
}
