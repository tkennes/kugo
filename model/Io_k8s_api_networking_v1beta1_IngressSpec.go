package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_Ingress.go


// IngressSpec describes the Ingress the user wishes to exist.
type Io_k8s_api_networking_v1beta1_IngressSpec struct {
	// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be
	// specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressBackend.go
	Backend          *Io_k8s_api_networking_v1beta1_IngressBackend `json:"backend,omitempty"`

	// IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller
	// will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards
	// compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a
	// warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a
	// class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this
	// field. For more information, refer to the IngressClass documentation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	IngressClassName *string                                       `json:"ingressClassName,omitempty"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the
	// default backend.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressRule.go
	Rules            []Io_k8s_api_networking_v1beta1_IngressRule   `json:"rules,omitempty"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify
	// different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS
	// extension, if the ingress controller fulfilling the ingress supports SNI.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressTLS.go
	Tls              []Io_k8s_api_networking_v1beta1_IngressTLS    `json:"tls,omitempty"`
}
