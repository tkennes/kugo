package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Namespace.go


// NamespaceSpec describes the attributes on a Namespace.
type Io_k8s_api_core_v1_NamespaceSpec struct {
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info:
	// https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Finalizers []*string `json:"finalizers,omitempty"`
}
