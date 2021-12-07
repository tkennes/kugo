package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
type Io_k8s_api_core_v1_FlexVolumeSource struct {
	// Driver is the name of the driver to use for this volume.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Driver    *string                                  `json:"driver"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// The default filesystem depends on FlexVolume script.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType    *string                                  `json:"fsType,omitempty"`

	// Optional: Extra command options if any.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Options   *interface{}                             `json:"options,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly  *bool                                    `json:"readOnly,omitempty"`

	// Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts.
	// This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are
	// passed to the plugin scripts.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LocalObjectReference.go
	SecretRef *Io_k8s_api_core_v1_LocalObjectReference `json:"secretRef,omitempty"`
}
