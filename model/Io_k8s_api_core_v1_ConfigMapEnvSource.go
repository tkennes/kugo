package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EnvFromSource.go


// ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.  The contents of the target
// ConfigMap's Data field will represent the key-value pairs as environment variables.
type Io_k8s_api_core_v1_ConfigMapEnvSource struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name     *string `json:"name,omitempty"`

	// Specify whether the ConfigMap must be defined
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Optional *bool   `json:"optional,omitempty"`
}
