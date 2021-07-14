package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecurityContext.go


// Adds and removes POSIX capabilities from running containers.
type Io_k8s_api_core_v1_Capabilities struct {
	// Added capabilities
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Add  []*string `json:"add,omitempty"`

	// Removed capabilities
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Drop []*string `json:"drop,omitempty"`
}
