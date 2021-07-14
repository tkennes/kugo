package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_IngressSpec.go


// IngressTLS describes the transport layer security associated with an Ingress.
type Io_k8s_api_extensions_v1beta1_IngressTLS struct {
	// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the
	// tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left
	// unspecified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Hosts      []*string `json:"hosts,omitempty"`

	// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing
	// based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an
	// IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SecretName *string   `json:"secretName,omitempty"`
}
