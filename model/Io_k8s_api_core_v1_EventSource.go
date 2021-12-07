package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_events_v1_Event.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_events_v1beta1_Event.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Event.go


// EventSource contains information for an event.
type Io_k8s_api_core_v1_EventSource struct {
	// Component from which the event is generated.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Component *string `json:"component,omitempty"`

	// Node name on which the event is generated.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Host      *string `json:"host,omitempty"`
}
