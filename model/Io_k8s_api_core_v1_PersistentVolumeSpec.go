package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolume.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachmentSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_VolumeAttachmentSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1alpha1_VolumeAttachmentSource.go


// PersistentVolumeSpec is the specification of a persistent volume.
type Io_k8s_api_core_v1_PersistentVolumeSpec struct {
	// AccessModes contains all ways the volume can be mounted. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	AccessModes                   []*string                                            `json:"accessModes,omitempty"`

	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to
	// the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource.go
	AwsElasticBlockStore          *Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_AzureDiskVolumeSource.go
	AzureDisk                     *Io_k8s_api_core_v1_AzureDiskVolumeSource            `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_AzureFilePersistentVolumeSource.go
	AzureFile                     *Io_k8s_api_core_v1_AzureFilePersistentVolumeSource  `json:"azureFile,omitempty"`

	// A description of the persistent volume's resources and capacity. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Capacity                      *interface{}                                         `json:"capacity,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CephFSPersistentVolumeSource.go
	Cephfs                        *Io_k8s_api_core_v1_CephFSPersistentVolumeSource     `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine. More info:
	// https://examples.k8s.io/mysql-cinder-pd/README.md
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CinderPersistentVolumeSource.go
	Cinder                        *Io_k8s_api_core_v1_CinderPersistentVolumeSource     `json:"cinder,omitempty"`

	// ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil
	// when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	ClaimRef                      *Io_k8s_api_core_v1_ObjectReference                  `json:"claimRef,omitempty"`

	// CSI represents storage that is handled by an external CSI driver (Beta feature).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CSIPersistentVolumeSource.go
	Csi                           *Io_k8s_api_core_v1_CSIPersistentVolumeSource        `json:"csi,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_FCVolumeSource.go
	Fc                            *Io_k8s_api_core_v1_FCVolumeSource                   `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_FlexPersistentVolumeSource.go
	FlexVolume                    *Io_k8s_api_core_v1_FlexPersistentVolumeSource       `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This
	// depends on the Flocker control service being running
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_FlockerVolumeSource.go
	Flocker                       *Io_k8s_api_core_v1_FlockerVolumeSource              `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the
	// pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource.go
	GcePersistentDisk             *Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource    `json:"gcePersistentDisk,omitempty"`

	// Glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More
	// info: https://examples.k8s.io/volumes/glusterfs/README.md
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_GlusterfsPersistentVolumeSource.go
	Glusterfs                     *Io_k8s_api_core_v1_GlusterfsPersistentVolumeSource  `json:"glusterfs,omitempty"`

	// HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node
	// development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_HostPathVolumeSource.go
	HostPath                      *Io_k8s_api_core_v1_HostPathVolumeSource             `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	// Provisioned by an admin.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ISCSIPersistentVolumeSource.go
	Iscsi                         *Io_k8s_api_core_v1_ISCSIPersistentVolumeSource      `json:"iscsi,omitempty"`

	// Local represents directly-attached storage with node affinity
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LocalVolumeSource.go
	Local                         *Io_k8s_api_core_v1_LocalVolumeSource                `json:"local,omitempty"`

	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MountOptions                  []*string                                            `json:"mountOptions,omitempty"`

	// NFS represents an NFS mount on the host. Provisioned by an admin. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NFSVolumeSource.go
	Nfs                           *Io_k8s_api_core_v1_NFSVolumeSource                  `json:"nfs,omitempty"`

	// NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the
	// scheduling of pods that use this volume.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VolumeNodeAffinity.go
	NodeAffinity                  *Io_k8s_api_core_v1_VolumeNodeAffinity               `json:"nodeAffinity,omitempty"`

	// What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created
	// PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle
	// must be supported by the volume plugin underlying this PersistentVolume. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	PersistentVolumeReclaimPolicy *string                                              `json:"persistentVolumeReclaimPolicy,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource.go
	PhotonPersistentDisk          *Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PortworxVolumeSource.go
	PortworxVolume                *Io_k8s_api_core_v1_PortworxVolumeSource             `json:"portworxVolume,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_QuobyteVolumeSource.go
	Quobyte                       *Io_k8s_api_core_v1_QuobyteVolumeSource              `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info:
	// https://examples.k8s.io/volumes/rbd/README.md
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_RBDPersistentVolumeSource.go
	Rbd                           *Io_k8s_api_core_v1_RBDPersistentVolumeSource        `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ScaleIOPersistentVolumeSource.go
	ScaleIO                       *Io_k8s_api_core_v1_ScaleIOPersistentVolumeSource    `json:"scaleIO,omitempty"`

	// Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any
	// StorageClass.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	StorageClassName              *string                                              `json:"storageClassName,omitempty"`

	// StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod More
	// info: https://examples.k8s.io/volumes/storageos/README.md
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_StorageOSPersistentVolumeSource.go
	Storageos                     *Io_k8s_api_core_v1_StorageOSPersistentVolumeSource  `json:"storageos,omitempty"`

	// volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value
	// of Filesystem is implied when not included in spec.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	VolumeMode                    *string                                              `json:"volumeMode,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource.go
	VsphereVolume                 *Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource   `json:"vsphereVolume,omitempty"`
}
