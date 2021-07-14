package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressRule.go


// HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart>
// -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything
// after the last '/' and before the first '?' or '#'.
type Io_k8s_api_networking_v1beta1_HTTPIngressRuleValue struct {
	// A collection of paths that map requests to backends.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_HTTPIngressPath.go
	Paths []Io_k8s_api_networking_v1beta1_HTTPIngressPath `json:"paths"`
}
