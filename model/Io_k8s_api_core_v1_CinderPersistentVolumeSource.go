package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume
// must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
type Io_k8s_api_core_v1_CinderPersistentVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs",
	// "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FsType    *string                             `json:"fsType,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info:
	// https://examples.k8s.io/mysql-cinder-pd/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly  *bool                               `json:"readOnly,omitempty"`

	// Optional: points to a secret object containing parameters used to connect to OpenStack.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretReference.go
	SecretRef *Io_k8s_api_core_v1_SecretReference `json:"secretRef,omitempty"`

	// volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	VolumeID  *string                             `json:"volumeID"`
}
