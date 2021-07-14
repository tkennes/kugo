package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EnvFromSource.go


// SecretEnvSource selects a Secret to populate the environment variables with.  The contents of the target Secret's Data
// field will represent the key-value pairs as environment variables.
type Io_k8s_api_core_v1_SecretEnvSource struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name     *string `json:"name,omitempty"`

	// Specify whether the Secret must be defined
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Optional *bool   `json:"optional,omitempty"`
}
