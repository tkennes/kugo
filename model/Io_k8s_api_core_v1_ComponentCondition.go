package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ComponentStatus.go


// Information about the condition of a component.
type Io_k8s_api_core_v1_ComponentCondition struct {
	// Condition error code for a component. For example, a health check error code.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Error   *string `json:"error,omitempty"`

	// Message about the condition for a component. For example, information about a health check.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message *string `json:"message,omitempty"`

	// Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status  *string `json:"status"`

	// Type of condition for a component. Valid value: "Healthy"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type    *string `json:"type"`
}
