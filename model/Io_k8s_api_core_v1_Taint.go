package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeSpec.go


// The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
type Io_k8s_api_core_v1_Taint struct {
	// Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule
	// and NoExecute.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Effect    *string    `json:"effect"`

	// Required. The taint key to be applied to a node.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Key       *string    `json:"key"`

	// TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	TimeAdded *time.Time `json:"timeAdded,omitempty"`

	// The taint value corresponding to the taint key.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Value     *string    `json:"value,omitempty"`
}
