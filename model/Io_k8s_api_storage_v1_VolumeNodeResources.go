package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSINodeDriver.go


// VolumeNodeResources is a set of resource limits for scheduling of volumes.
type Io_k8s_api_storage_v1_VolumeNodeResources struct {
	// Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached
	// and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared
	// among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node
	// is unbounded.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Count *int `json:"count,omitempty"`
}
