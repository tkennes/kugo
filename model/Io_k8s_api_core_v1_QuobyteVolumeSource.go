package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or
// SELinux relabeling.
type Io_k8s_api_core_v1_QuobyteVolumeSource struct {
	// Group to map volume access to Default is no group
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Group    *string `json:"group,omitempty"`

	// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly *bool   `json:"readOnly,omitempty"`

	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple
	// entries are separated with commas) which acts as the central registry for volumes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Registry *string `json:"registry"`

	// Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by
	// the plugin
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Tenant   *string `json:"tenant,omitempty"`

	// User to map volume access to Defaults to serivceaccount user
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	User     *string `json:"user,omitempty"`

	// Volume is a string that references an already created Quobyte volume by name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Volume   *string `json:"volume"`
}
