package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSecurityContext.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecurityContext.go


// SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
type Io_k8s_api_core_v1_SeccompProfile struct {
	// localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on
	// the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be
	// set if type is "Localhost".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	LocalhostProfile *string `json:"localhostProfile,omitempty"`

	// type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a
	// file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no
	// profile should be applied.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type             *string `json:"type"`
}
