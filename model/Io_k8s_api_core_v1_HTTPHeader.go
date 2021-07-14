package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_HTTPGetAction.go


// HTTPHeader describes a custom header to be used in HTTP probes
type Io_k8s_api_core_v1_HTTPHeader struct {
	// The header field name
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name  *string `json:"name"`

	// The header field value
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Value *string `json:"value"`
}
