package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Container.go


// EnvVar represents an environment variable present in a Container.
type Io_k8s_api_core_v1_EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name      *string                          `json:"name"`

	// Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any
	// service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged.
	// The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not. Defaults to "".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Value     *string                          `json:"value,omitempty"`

	// Source for the environment variable's value. Cannot be used if value is not empty.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EnvVarSource.go
	ValueFrom *Io_k8s_api_core_v1_EnvVarSource `json:"valueFrom,omitempty"`
}
