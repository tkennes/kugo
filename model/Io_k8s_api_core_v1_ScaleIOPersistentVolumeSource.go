package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
type Io_k8s_api_core_v1_ScaleIOPersistentVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// Default is "xfs"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType           *string                            `json:"fsType,omitempty"`

	// The host address of the ScaleIO API Gateway.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Gateway          *string                            `json:"gateway"`

	// The name of the ScaleIO Protection Domain for the configured storage.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ProtectionDomain *string                            `json:"protectionDomain,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly         *bool                              `json:"readOnly,omitempty"`

	// SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login
	// operation will fail.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	SecretRef        Io_k8s_api_core_v1_SecretReference `json:"secretRef"`

	// Flag to enable/disable SSL communication with Gateway, default false
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	SslEnabled       *bool                              `json:"sslEnabled,omitempty"`

	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	StorageMode      *string                            `json:"storageMode,omitempty"`

	// The ScaleIO Storage Pool associated with the protection domain.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	StoragePool      *string                            `json:"storagePool,omitempty"`

	// The name of the storage system as configured in ScaleIO.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	System           *string                            `json:"system"`

	// The name of a volume already created in the ScaleIO system that is associated with this volume source.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumeName       *string                            `json:"volumeName,omitempty"`
}
