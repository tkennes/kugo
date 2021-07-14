package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Photon Controller persistent disk resource.
type Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// Implicitly inferred to be "ext4" if unspecified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType *string `json:"fsType,omitempty"`

	// ID that identifies Photon Controller persistent disk
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	PdID   *string `json:"pdID"`
}
