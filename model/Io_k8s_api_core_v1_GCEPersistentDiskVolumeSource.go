package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Persistent Disk resource in Google Compute Engine.  A GCE PD must exist before mounting to a container. The
// disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or
// read-only many times. GCE PDs support ownership management and SELinux relabeling.
type Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host
	// operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType    *string `json:"fsType,omitempty"`

	// The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For
	// volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can
	// leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Partition *int    `json:"partition,omitempty"`

	// Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PdName    *string `json:"pdName"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}
