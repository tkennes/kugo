package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_Ingress.go


// IngressSpec describes the Ingress the user wishes to exist.
type Io_k8s_api_networking_v1_IngressSpec struct {
	// DefaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified,
	// DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the
	// rules will be up to the Ingress controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressBackend.go
	DefaultBackend   *Io_k8s_api_networking_v1_IngressBackend `json:"defaultBackend,omitempty"`

	// IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller
	// will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards
	// compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a
	// warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a
	// class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this
	// field. For more information, refer to the IngressClass documentation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	IngressClassName *string                                  `json:"ingressClassName,omitempty"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the
	// default backend.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressRule.go
	Rules            []Io_k8s_api_networking_v1_IngressRule   `json:"rules,omitempty"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify
	// different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS
	// extension, if the ingress controller fulfilling the ingress supports SNI.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressTLS.go
	Tls              []Io_k8s_api_networking_v1_IngressTLS    `json:"tls,omitempty"`
}
