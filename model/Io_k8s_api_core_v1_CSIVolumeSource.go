package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a source location of a volume to mount, managed by an external CSI driver
type Io_k8s_api_core_v1_CSIVolumeSource struct {
	// Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as
	// registered in the cluster.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Driver               *string                                  `json:"driver"`

	// Filesystem type to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI
	// driver which will determine the default filesystem to apply.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType               *string                                  `json:"fsType,omitempty"`

	// NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to
	// complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret
	// is required. If the secret object contains more than one secret, all secret references are passed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LocalObjectReference.go
	NodePublishSecretRef *Io_k8s_api_core_v1_LocalObjectReference `json:"nodePublishSecretRef,omitempty"`

	// Specifies a read-only configuration for the volume. Defaults to false (read/write).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly             *bool                                    `json:"readOnly,omitempty"`

	// VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's
	// documentation for supported values.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	VolumeAttributes     *interface{}                             `json:"volumeAttributes,omitempty"`
}
