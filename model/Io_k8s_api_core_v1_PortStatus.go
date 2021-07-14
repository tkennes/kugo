package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LoadBalancerIngress.go
type Io_k8s_api_core_v1_PortStatus struct {
	// Error is to record the problem with the service port The format of the error shall comply with the following rules: -
	// built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific
	// error values must have names that comply with the   format foo.example.com/CamelCase.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Error    *string `json:"error,omitempty"`

	// Port is the port number of the service port of which status is recorded here
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Port     *int    `json:"port"`

	// Protocol is the protocol of the service port of which status is recorded here The supported values are: "TCP", "UDP",
	// "SCTP"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Protocol *string `json:"protocol"`
}
