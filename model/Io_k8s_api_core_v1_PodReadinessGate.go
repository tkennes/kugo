package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSpec.go


// PodReadinessGate contains the reference to a pod condition
type Io_k8s_api_core_v1_PodReadinessGate struct {
	// ConditionType refers to a condition in the pod's condition list with matching type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ConditionType *string `json:"conditionType"`
}
