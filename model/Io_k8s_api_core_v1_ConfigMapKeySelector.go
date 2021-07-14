package kugo_model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EnvVarSource.go


// Selects a key from a ConfigMap.
type Io_k8s_api_core_v1_ConfigMapKeySelector struct {
	// The key to select.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Key      *string `json:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name     *string `json:"name,omitempty"`

	// Specify whether the ConfigMap or its key must be defined
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Optional *bool   `json:"optional,omitempty"`
}
