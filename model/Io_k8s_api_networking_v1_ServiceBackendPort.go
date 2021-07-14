package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressServiceBackend.go


// ServiceBackendPort is the service port being referenced.
type Io_k8s_api_networking_v1_ServiceBackendPort struct {
	// Name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name   *string `json:"name,omitempty"`

	// Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Number *int    `json:"number,omitempty"`
}
