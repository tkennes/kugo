package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1_RuntimeClass.go


// Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
type Io_k8s_api_node_v1_Scheduling struct {
	// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can
	// only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing
	// nodeSelector. Any conflicts will cause the pod to be rejected in admission.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	NodeSelector *interface{}                    `json:"nodeSelector,omitempty"`

	// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively
	// unioning the set of nodes tolerated by the pod and the RuntimeClass.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Toleration.go
	Tolerations  []Io_k8s_api_core_v1_Toleration `json:"tolerations,omitempty"`
}
