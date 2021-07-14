package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_StatefulSetSpec.go


// StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It
// includes any additional parameters necessary to perform the update for the indicated strategy.
type Io_k8s_api_apps_v1_StatefulSetUpdateStrategy struct {
	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_RollingUpdateStatefulSetStrategy.go
	RollingUpdate *Io_k8s_api_apps_v1_RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`

	// Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type          *string                                              `json:"type,omitempty"`
}
