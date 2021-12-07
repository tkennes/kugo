package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec.go


// AllowedHostPath defines the host volume conditions that will be enabled by a policy for pods to use. It requires the
// path prefix to be defined.
type Io_k8s_api_policy_v1beta1_AllowedHostPath struct {
	// pathPrefix is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed
	// when validating the path prefix with a host path.  Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo`
	// would not allow `/food` or `/etc/foo`
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PathPrefix *string `json:"pathPrefix,omitempty"`

	// when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly   *bool   `json:"readOnly,omitempty"`
}
