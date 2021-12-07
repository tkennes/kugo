package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and
// SELinux relabeling.
type Io_k8s_api_core_v1_RBDPersistentVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host
	// operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#rbd
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType    *string                             `json:"fsType,omitempty"`

	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Image     *string                             `json:"image"`

	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Keyring   *string                             `json:"keyring,omitempty"`

	// A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Monitors  []*string                           `json:"monitors"`

	// The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Pool      *string                             `json:"pool,omitempty"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly  *bool                               `json:"readOnly,omitempty"`

	// SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	SecretRef *Io_k8s_api_core_v1_SecretReference `json:"secretRef,omitempty"`

	// The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	User      *string                             `json:"user,omitempty"`
}
