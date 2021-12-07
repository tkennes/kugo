package kugo_model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EnvVarSource.go


// SecretKeySelector selects a key of a Secret.
type Io_k8s_api_core_v1_SecretKeySelector struct {
	// The key of the secret to select from.  Must be a valid secret key.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Key      *string `json:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name     *string `json:"name,omitempty"`

	// Specify whether the Secret or its key must be defined
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Optional *bool   `json:"optional,omitempty"`
}
