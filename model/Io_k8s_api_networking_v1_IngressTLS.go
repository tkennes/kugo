package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressSpec.go


// IngressTLS describes the transport layer security associated with an Ingress.
type Io_k8s_api_networking_v1_IngressTLS struct {
	// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the
	// tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left
	// unspecified.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Hosts      []*string `json:"hosts,omitempty"`

	// SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS
	// routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an
	// IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	SecretName *string   `json:"secretName,omitempty"`
}
