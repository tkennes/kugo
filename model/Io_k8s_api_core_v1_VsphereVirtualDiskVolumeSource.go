package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a vSphere volume resource.
type Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// Implicitly inferred to be "ext4" if unspecified.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType            *string `json:"fsType,omitempty"`

	// Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	StoragePolicyID   *string `json:"storagePolicyID,omitempty"`

	// Storage Policy Based Management (SPBM) profile name.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	StoragePolicyName *string `json:"storagePolicyName,omitempty"`

	// Path that identifies vSphere volume vmdk
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	VolumePath        *string `json:"volumePath"`
}
