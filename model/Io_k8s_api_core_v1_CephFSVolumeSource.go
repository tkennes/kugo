package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management
// or SELinux relabeling.
type Io_k8s_api_core_v1_CephFSVolumeSource struct {
	// Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-
	// use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Monitors   []*string                                `json:"monitors"`

	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Path       *string                                  `json:"path,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info:
	// https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly   *bool                                    `json:"readOnly,omitempty"`

	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info:
	// https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	SecretFile *string                                  `json:"secretFile,omitempty"`

	// Optional: SecretRef is reference to the authentication secret for User, default is empty. More info:
	// https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LocalObjectReference.go
	SecretRef  *Io_k8s_api_core_v1_LocalObjectReference `json:"secretRef,omitempty"`

	// Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-
	// to-use-it
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	User       *string                                  `json:"user,omitempty"`
}
