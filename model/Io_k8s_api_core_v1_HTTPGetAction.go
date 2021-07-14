package model


// HTTPGetAction describes an action based on HTTP Get requests.
type Io_k8s_api_core_v1_HTTPGetAction struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Host        *string                         `json:"host,omitempty"`

	// Custom headers to set in the request. HTTP allows repeated headers.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_HTTPHeader.go
	HttpHeaders []Io_k8s_api_core_v1_HTTPHeader `json:"httpHeaders,omitempty"`

	// Path to access on the HTTP server.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path        *string                         `json:"path,omitempty"`

	// Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an
	// IANA_SVC_NAME.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Port        *int                            `json:"port"`

	// Scheme to use for connecting to the host. Defaults to HTTP.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Scheme      *string                         `json:"scheme,omitempty"`
}
