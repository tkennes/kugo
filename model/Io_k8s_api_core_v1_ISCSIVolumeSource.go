package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership
// management and SELinux relabeling.
type Io_k8s_api_core_v1_ISCSIVolumeSource struct {
	// whether support iSCSI Discovery CHAP authentication
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ChapAuthDiscovery *bool                                    `json:"chapAuthDiscovery,omitempty"`

	// whether support iSCSI Session CHAP authentication
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ChapAuthSession   *bool                                    `json:"chapAuthSession,omitempty"`

	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host
	// operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FsType            *string                                  `json:"fsType,omitempty"`

	// Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
	// <target portal>:<volume name> will be created for the connection.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	InitiatorName     *string                                  `json:"initiatorName,omitempty"`

	// Target iSCSI Qualified Name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Iqn               *string                                  `json:"iqn"`

	// iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	IscsiInterface    *string                                  `json:"iscsiInterface,omitempty"`

	// iSCSI Target Lun number.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Lun               *int                                     `json:"lun"`

	// iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP
	// ports 860 and 3260).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Portals           []*string                                `json:"portals,omitempty"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly          *bool                                    `json:"readOnly,omitempty"`

	// CHAP Secret for iSCSI target and initiator authentication
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LocalObjectReference.go
	SecretRef         *Io_k8s_api_core_v1_LocalObjectReference `json:"secretRef,omitempty"`

	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports
	// 860 and 3260).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	TargetPortal      *string                                  `json:"targetPortal"`
}
