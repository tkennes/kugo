package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Node.go


// NodeSpec describes the attributes that a node is created with.
type Io_k8s_api_core_v1_NodeSpec struct {
	// If specified, the source to get node configuration from The DynamicKubeletConfig feature gate must be enabled for the
	// Kubelet to use this field
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeConfigSource.go
	ConfigSource  *Io_k8s_api_core_v1_NodeConfigSource `json:"configSource,omitempty"`

	// Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ExternalID    *string                              `json:"externalID,omitempty"`

	// PodCIDR represents the pod IP range assigned to the node.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PodCIDR       *string                              `json:"podCIDR,omitempty"`

	// podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the
	// 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PodCIDRs      []*string                            `json:"podCIDRs,omitempty"`

	// ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ProviderID    *string                              `json:"providerID,omitempty"`

	// If specified, the node's taints.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Taint.go
	Taints        []Io_k8s_api_core_v1_Taint           `json:"taints,omitempty"`

	// Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Unschedulable *bool                                `json:"unschedulable,omitempty"`
}
