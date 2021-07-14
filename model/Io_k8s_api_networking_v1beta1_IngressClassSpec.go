package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressClass.go


// IngressClassSpec provides information about the class of an Ingress.
type Io_k8s_api_networking_v1beta1_IngressClassSpec struct {
	// Controller refers to the name of the controller that should handle this class. This allows for different "flavors" that
	// are controlled by the same controller. For example, you may have different Parameters for the same implementing
	// controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g.
	// "acme.io/ingress-controller". This field is immutable.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Controller *string                                       `json:"controller,omitempty"`

	// Parameters is a link to a custom resource containing additional configuration for the controller. This is optional if
	// the controller does not require extra parameters.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_TypedLocalObjectReference.go
	Parameters *Io_k8s_api_core_v1_TypedLocalObjectReference `json:"parameters,omitempty"`
}
