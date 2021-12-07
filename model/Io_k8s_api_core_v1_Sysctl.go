package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodSecurityContext.go


// Sysctl defines a kernel parameter to be set
type Io_k8s_api_core_v1_Sysctl struct {
	// Name of a property to set
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name  *string `json:"name"`

	// Value of a property to set
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Value *string `json:"value"`
}
