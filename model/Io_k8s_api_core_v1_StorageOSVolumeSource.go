package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a StorageOS persistent volume resource.
type Io_k8s_api_core_v1_StorageOSVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// Implicitly inferred to be "ext4" if unspecified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType          *string                                  `json:"fsType,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly        *bool                                    `json:"readOnly,omitempty"`

	// SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values
	// will be attempted.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LocalObjectReference.go
	SecretRef       *Io_k8s_api_core_v1_LocalObjectReference `json:"secretRef,omitempty"`

	// VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumeName      *string                                  `json:"volumeName,omitempty"`

	// VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's
	// namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter
	// integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using
	// namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumeNamespace *string                                  `json:"volumeNamespace,omitempty"`
}
