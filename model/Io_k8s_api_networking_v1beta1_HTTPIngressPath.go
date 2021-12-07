package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1beta1_HTTPIngressRuleValue.go


// HTTPIngressPath associates a path with a backend. Incoming urls matching the path are forwarded to the backend.
type Io_k8s_api_networking_v1beta1_HTTPIngressPath struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1beta1_IngressBackend.go
	Backend  Io_k8s_api_networking_v1beta1_IngressBackend `json:"backend"`

	// Path is matched against the path of an incoming request. Currently it can contain characters disallowed from the
	// conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. When unspecified, all paths from
	// incoming requests are matched.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Path     *string                                      `json:"path,omitempty"`

	// PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact:
	// Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is   done on a path
	// element by element basis. A path element refers is the   list of labels in the path split by the '/' separator. A
	// request is a   match for path p if every p is an element-wise prefix of p of the   request path. Note that if the last
	// element of the path is a substring   of the last element in request path, it is not a match (e.g. /foo/bar   matches
	// /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to
	// the IngressClass. Implementations can treat this as a separate PathType   or treat it identically to Prefix or Exact
	// path types. Implementations are required to support all path types. Defaults to ImplementationSpecific.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PathType *string                                      `json:"pathType,omitempty"`
}
