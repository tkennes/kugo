package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes
// support ownership management and SELinux relabeling.
type Io_k8s_api_core_v1_FCVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// Implicitly inferred to be "ext4" if unspecified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType     *string   `json:"fsType,omitempty"`

	// Optional: FC target lun number
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Lun        *int      `json:"lun,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly   *bool     `json:"readOnly,omitempty"`

	// Optional: FC target worldwide names (WWNs)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	TargetWWNs []*string `json:"targetWWNs,omitempty"`

	// Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but
	// not both simultaneously.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Wwids      []*string `json:"wwids,omitempty"`
}
