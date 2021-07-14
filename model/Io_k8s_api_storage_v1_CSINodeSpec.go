package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSINode.go


// CSINodeSpec holds information about the specification of all CSI drivers installed on a node
type Io_k8s_api_storage_v1_CSINodeSpec struct {
	// drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this
	// can become empty.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSINodeDriver.go
	Drivers []Io_k8s_api_storage_v1_CSINodeDriver `json:"drivers"`
}
