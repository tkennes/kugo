package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// PortworxVolumeSource represents a Portworx volume resource.
type Io_k8s_api_core_v1_PortworxVolumeSource struct {
	// FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex.
	// "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType   *string `json:"fsType,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly *bool   `json:"readOnly,omitempty"`

	// VolumeID uniquely identifies a Portworx volume
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumeID *string `json:"volumeID"`
}
