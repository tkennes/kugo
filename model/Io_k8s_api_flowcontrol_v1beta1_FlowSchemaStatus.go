package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_FlowSchema.go


// FlowSchemaStatus represents the current state of a FlowSchema.
type Io_k8s_api_flowcontrol_v1beta1_FlowSchemaStatus struct {
	// `conditions` is a list of the current states of FlowSchema.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_FlowSchemaCondition.go
	Conditions []Io_k8s_api_flowcontrol_v1beta1_FlowSchemaCondition `json:"conditions,omitempty"`
}
