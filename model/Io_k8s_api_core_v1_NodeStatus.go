package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Node.go


// NodeStatus is information about the current status of a node.
type Io_k8s_api_core_v1_NodeStatus struct {
	// List of addresses reachable to the node. Queried from cloud provider, if available. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key
	// is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-
	// replacement patch. See http://pr.k8s.io/79391 for an example.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeAddress.go
	Addresses       []Io_k8s_api_core_v1_NodeAddress        `json:"addresses,omitempty"`

	// Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Allocatable     *interface{}                            `json:"allocatable,omitempty"`

	// Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-
	// volumes#capacity
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Capacity        *interface{}                            `json:"capacity,omitempty"`

	// Conditions is an array of current observed node conditions. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#condition
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeCondition.go
	Conditions      []Io_k8s_api_core_v1_NodeCondition      `json:"conditions,omitempty"`

	// Status of the config assigned to the node via the dynamic Kubelet config feature.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeConfigStatus.go
	Config          *Io_k8s_api_core_v1_NodeConfigStatus    `json:"config,omitempty"`

	// Endpoints of daemons running on the Node.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeDaemonEndpoints.go
	DaemonEndpoints *Io_k8s_api_core_v1_NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`

	// List of container images on this node
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerImage.go
	Images          []Io_k8s_api_core_v1_ContainerImage     `json:"images,omitempty"`

	// Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#info
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSystemInfo.go
	NodeInfo        *Io_k8s_api_core_v1_NodeSystemInfo      `json:"nodeInfo,omitempty"`

	// NodePhase is the recently observed lifecycle phase of the node. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Phase           *string                                 `json:"phase,omitempty"`

	// List of volumes that are attached to the node.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_AttachedVolume.go
	VolumesAttached []Io_k8s_api_core_v1_AttachedVolume     `json:"volumesAttached,omitempty"`

	// List of attachable volumes in use (mounted) by the node.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumesInUse    []*string                               `json:"volumesInUse,omitempty"`
}
