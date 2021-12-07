package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecurityContext.go


// Adds and removes POSIX capabilities from running containers.
type Io_k8s_api_core_v1_Capabilities struct {
	// Added capabilities
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Add  []*string `json:"add,omitempty"`

	// Removed capabilities
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Drop []*string `json:"drop,omitempty"`
}
