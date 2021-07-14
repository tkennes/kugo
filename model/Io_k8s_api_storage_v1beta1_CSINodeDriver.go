package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_CSINodeSpec.go


// CSINodeDriver holds information about the specification of one CSI driver installed on a node
type Io_k8s_api_storage_v1beta1_CSINodeDriver struct {
	// allocatable represents the volume resources of a node that are available for scheduling.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_VolumeNodeResources.go
	Allocatable  *Io_k8s_api_storage_v1beta1_VolumeNodeResources `json:"allocatable,omitempty"`

	// This is the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI
	// GetPluginName() call for that driver.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name         *string                                         `json:"name"`

	// nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that
	// do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the
	// storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a
	// volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will
	// understand, e.g. "nodeA" instead of "node1". This field is required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	NodeID       *string                                         `json:"nodeID"`

	// topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set
	// of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a
	// node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own
	// node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should
	// retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology
	// keys. This can be empty if driver does not support topology.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	TopologyKeys []*string                                       `json:"topologyKeys,omitempty"`
}
