package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSecurityContext.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_SELinuxStrategyOptions.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecurityContext.go


// SELinuxOptions are the labels to be applied to the container
type Io_k8s_api_core_v1_SELinuxOptions struct {
	// Level is SELinux level label that applies to the container.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Level *string `json:"level,omitempty"`

	// Role is a SELinux role label that applies to the container.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Role  *string `json:"role,omitempty"`

	// Type is a SELinux type label that applies to the container.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type  *string `json:"type,omitempty"`

	// User is a SELinux user label that applies to the container.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	User  *string `json:"user,omitempty"`
}
