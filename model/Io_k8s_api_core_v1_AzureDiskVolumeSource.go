package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type Io_k8s_api_core_v1_AzureDiskVolumeSource struct {
	// Host Caching mode: None, Read Only, Read Write.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	CachingMode *string `json:"cachingMode,omitempty"`

	// The Name of the data disk in the blob storage
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DiskName    *string `json:"diskName"`

	// The URI the data disk in the blob storage
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DiskURI     *string `json:"diskURI"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// Implicitly inferred to be "ext4" if unspecified.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType      *string `json:"fsType,omitempty"`

	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account
	// Managed: azure managed data disk (only in managed availability set). defaults to shared
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind        *string `json:"kind,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly    *bool   `json:"readOnly,omitempty"`
}
