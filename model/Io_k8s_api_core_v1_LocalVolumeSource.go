package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// Local represents directly-attached storage with node affinity (Beta feature)
type Io_k8s_api_core_v1_LocalVolumeSource struct {
	// Filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the
	// host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a fileystem if unspecified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType *string `json:"fsType,omitempty"`

	// The full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path   *string `json:"path"`
}
