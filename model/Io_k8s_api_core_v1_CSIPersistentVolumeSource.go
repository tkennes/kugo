package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// Represents storage that is managed by an external CSI volume driver (Beta feature)
type Io_k8s_api_core_v1_CSIPersistentVolumeSource struct {
	// ControllerExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver
	// to complete the CSI ControllerExpandVolume call. This is an alpha field and requires enabling ExpandCSIVolumes feature
	// gate. This field is optional, and may be empty if no secret is required. If the secret object contains more than one
	// secret, all secrets are passed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	ControllerExpandSecretRef  *Io_k8s_api_core_v1_SecretReference `json:"controllerExpandSecretRef,omitempty"`

	// ControllerPublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI
	// driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. This field is optional, and may
	// be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	ControllerPublishSecretRef *Io_k8s_api_core_v1_SecretReference `json:"controllerPublishSecretRef,omitempty"`

	// Driver is the name of the driver to use for this volume. Required.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Driver                     *string                             `json:"driver"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType                     *string                             `json:"fsType,omitempty"`

	// NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to
	// complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret
	// is required. If the secret object contains more than one secret, all secrets are passed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	NodePublishSecretRef       *Io_k8s_api_core_v1_SecretReference `json:"nodePublishSecretRef,omitempty"`

	// NodeStageSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to
	// complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. This field is optional, and may be
	// empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	NodeStageSecretRef         *Io_k8s_api_core_v1_SecretReference `json:"nodeStageSecretRef,omitempty"`

	// Optional: The value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly                   *bool                               `json:"readOnly,omitempty"`

	// Attributes of the volume to publish.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	VolumeAttributes           *interface{}                        `json:"volumeAttributes,omitempty"`

	// VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all
	// subsequent calls. Required.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	VolumeHandle               *string                             `json:"volumeHandle"`
}
