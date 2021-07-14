package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Namespace.go


// NamespaceStatus is information about the current status of a Namespace.
type Io_k8s_api_core_v1_NamespaceStatus struct {
	// Represents the latest available observations of a namespace's current state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NamespaceCondition.go
	Conditions []Io_k8s_api_core_v1_NamespaceCondition `json:"conditions,omitempty"`

	// Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-
	// cluster/namespaces/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Phase      *string                                 `json:"phase,omitempty"`
}
