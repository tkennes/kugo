package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Container.go


// EnvFromSource represents the source of a set of ConfigMaps
type Io_k8s_api_core_v1_EnvFromSource struct {
	// The ConfigMap to select from
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMapEnvSource.go
	ConfigMapRef *Io_k8s_api_core_v1_ConfigMapEnvSource `json:"configMapRef,omitempty"`

	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Prefix       *string                                `json:"prefix,omitempty"`

	// The Secret to select from
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretEnvSource.go
	SecretRef    *Io_k8s_api_core_v1_SecretEnvSource    `json:"secretRef,omitempty"`
}
