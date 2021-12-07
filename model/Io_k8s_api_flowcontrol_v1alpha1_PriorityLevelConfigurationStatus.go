package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfiguration.go


// PriorityLevelConfigurationStatus represents the current state of a "request-priority".
type Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationStatus struct {
	// `conditions` is the current state of "request-priority".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationCondition.go
	Conditions []Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationCondition `json:"conditions,omitempty"`
}
