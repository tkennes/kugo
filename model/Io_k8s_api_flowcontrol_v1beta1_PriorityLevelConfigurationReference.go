package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_FlowSchemaSpec.go


// PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.
type Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationReference struct {
	// `name` is the name of the priority level configuration being referenced Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name *string `json:"name"`
}
