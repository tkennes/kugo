package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Persistent Disk resource in AWS.  An AWS EBS disk must exist before mounting to a container. The disk must
// also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes
// support ownership management and SELinux relabeling.
type Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host
	// operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType    *string `json:"fsType,omitempty"`

	// The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For
	// volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can
	// leave the property empty).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Partition *int    `json:"partition,omitempty"`

	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true". If omitted, the default is "false".
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly  *bool   `json:"readOnly,omitempty"`

	// Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumeID  *string `json:"volumeID"`
}
