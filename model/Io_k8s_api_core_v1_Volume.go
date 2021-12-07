package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodSpec.go


// Volume represents a named volume in a pod that may be accessed by any container in the pod.
type Io_k8s_api_core_v1_Volume struct {
	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to
	// the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource.go
	AwsElasticBlockStore  *Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource  `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_AzureDiskVolumeSource.go
	AzureDisk             *Io_k8s_api_core_v1_AzureDiskVolumeSource             `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_AzureFileVolumeSource.go
	AzureFile             *Io_k8s_api_core_v1_AzureFileVolumeSource             `json:"azureFile,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_CephFSVolumeSource.go
	Cephfs                *Io_k8s_api_core_v1_CephFSVolumeSource                `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine. More info:
	// https://examples.k8s.io/mysql-cinder-pd/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_CinderVolumeSource.go
	Cinder                *Io_k8s_api_core_v1_CinderVolumeSource                `json:"cinder,omitempty"`

	// ConfigMap represents a configMap that should populate this volume
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ConfigMapVolumeSource.go
	ConfigMap             *Io_k8s_api_core_v1_ConfigMapVolumeSource             `json:"configMap,omitempty"`

	// CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta
	// feature).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_CSIVolumeSource.go
	Csi                   *Io_k8s_api_core_v1_CSIVolumeSource                   `json:"csi,omitempty"`

	// DownwardAPI represents downward API about the pod that should populate this volume
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_DownwardAPIVolumeSource.go
	DownwardAPI           *Io_k8s_api_core_v1_DownwardAPIVolumeSource           `json:"downwardAPI,omitempty"`

	// EmptyDir represents a temporary directory that shares a pod's lifetime. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EmptyDirVolumeSource.go
	EmptyDir              *Io_k8s_api_core_v1_EmptyDirVolumeSource              `json:"emptyDir,omitempty"`

	// Ephemeral represents a volume that is handled by a cluster storage driver (Alpha feature). The volume's lifecycle is
	// tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use
	// this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or
	// capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver
	// supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more
	// information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one
	// of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for
	// light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the
	// driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EphemeralVolumeSource.go
	Ephemeral             *Io_k8s_api_core_v1_EphemeralVolumeSource             `json:"ephemeral,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_FCVolumeSource.go
	Fc                    *Io_k8s_api_core_v1_FCVolumeSource                    `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_FlexVolumeSource.go
	FlexVolume            *Io_k8s_api_core_v1_FlexVolumeSource                  `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service
	// being running
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_FlockerVolumeSource.go
	Flocker               *Io_k8s_api_core_v1_FlockerVolumeSource               `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the
	// pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource.go
	GcePersistentDisk     *Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource     `json:"gcePersistentDisk,omitempty"`

	// GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a
	// container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the
	// EmptyDir into the Pod's container.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_GitRepoVolumeSource.go
	GitRepo               *Io_k8s_api_core_v1_GitRepoVolumeSource               `json:"gitRepo,omitempty"`

	// Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_GlusterfsVolumeSource.go
	Glusterfs             *Io_k8s_api_core_v1_GlusterfsVolumeSource             `json:"glusterfs,omitempty"`

	// HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This
	// is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers
	// will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_HostPathVolumeSource.go
	HostPath              *Io_k8s_api_core_v1_HostPathVolumeSource              `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More
	// info: https://examples.k8s.io/volumes/iscsi/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ISCSIVolumeSource.go
	Iscsi                 *Io_k8s_api_core_v1_ISCSIVolumeSource                 `json:"iscsi,omitempty"`

	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name                  *string                                               `json:"name"`

	// NFS represents an NFS mount on the host that shares a pod's lifetime More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NFSVolumeSource.go
	Nfs                   *Io_k8s_api_core_v1_NFSVolumeSource                   `json:"nfs,omitempty"`

	// PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimVolumeSource.go
	PersistentVolumeClaim *Io_k8s_api_core_v1_PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource.go
	PhotonPersistentDisk  *Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource  `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PortworxVolumeSource.go
	PortworxVolume        *Io_k8s_api_core_v1_PortworxVolumeSource              `json:"portworxVolume,omitempty"`

	// Items for all in one resources secrets, configmaps, and downward API
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ProjectedVolumeSource.go
	Projected             *Io_k8s_api_core_v1_ProjectedVolumeSource             `json:"projected,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_QuobyteVolumeSource.go
	Quobyte               *Io_k8s_api_core_v1_QuobyteVolumeSource               `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info:
	// https://examples.k8s.io/volumes/rbd/README.md
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_RBDVolumeSource.go
	Rbd                   *Io_k8s_api_core_v1_RBDVolumeSource                   `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ScaleIOVolumeSource.go
	ScaleIO               *Io_k8s_api_core_v1_ScaleIOVolumeSource               `json:"scaleIO,omitempty"`

	// Secret represents a secret that should populate this volume. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#secret
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretVolumeSource.go
	Secret                *Io_k8s_api_core_v1_SecretVolumeSource                `json:"secret,omitempty"`

	// StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_StorageOSVolumeSource.go
	Storageos             *Io_k8s_api_core_v1_StorageOSVolumeSource             `json:"storageos,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource.go
	VsphereVolume         *Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource    `json:"vsphereVolume,omitempty"`
}
