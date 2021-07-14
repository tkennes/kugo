package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1alpha1_RuntimeClass.go


// RuntimeClassSpec is a specification of a RuntimeClass. It contains parameters that are required to describe the
// RuntimeClass to the Container Runtime Interface (CRI) implementation, as well as any other components that need to
// understand how the pod will be run. The RuntimeClassSpec is immutable.
type Io_k8s_api_node_v1alpha1_RuntimeClassSpec struct {
	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15,
	// and is only honored by servers that enable the PodOverhead feature.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1alpha1_Overhead.go
	Overhead       *Io_k8s_api_node_v1alpha1_Overhead   `json:"overhead,omitempty"`

	// RuntimeHandler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of
	// this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are
	// available on every node, and handlers of the same name are equivalent on every node. For example, a handler called
	// "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a
	// pod. The RuntimeHandler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	RuntimeHandler *string                              `json:"runtimeHandler"`

	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes
	// that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1alpha1_Scheduling.go
	Scheduling     *Io_k8s_api_node_v1alpha1_Scheduling `json:"scheduling,omitempty"`
}
