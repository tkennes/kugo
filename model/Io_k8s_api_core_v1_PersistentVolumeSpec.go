package kugo_model

type Io_k8s_api_core_v1_PersistentVolumeSpec struct {
	AccessModes                   []string                                             `json:"accessModes,omitempty"`
	AwsElasticBlockStore          *Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`
	AzureDisk                     *Io_k8s_api_core_v1_AzureDiskVolumeSource            `json:"azureDisk,omitempty"`
	AzureFile                     *Io_k8s_api_core_v1_AzureFilePersistentVolumeSource  `json:"azureFile,omitempty"`
	Capacity                      *interface{}                                         `json:"capacity,omitempty"`
	Cephfs                        *Io_k8s_api_core_v1_CephFSPersistentVolumeSource     `json:"cephfs,omitempty"`
	Cinder                        *Io_k8s_api_core_v1_CinderPersistentVolumeSource     `json:"cinder,omitempty"`
	ClaimRef                      *Io_k8s_api_core_v1_ObjectReference                  `json:"claimRef,omitempty"`
	Csi                           *Io_k8s_api_core_v1_CSIPersistentVolumeSource        `json:"csi,omitempty"`
	Fc                            *Io_k8s_api_core_v1_FCVolumeSource                   `json:"fc,omitempty"`
	FlexVolume                    *Io_k8s_api_core_v1_FlexPersistentVolumeSource       `json:"flexVolume,omitempty"`
	Flocker                       *Io_k8s_api_core_v1_FlockerVolumeSource              `json:"flocker,omitempty"`
	GcePersistentDisk             *Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource    `json:"gcePersistentDisk,omitempty"`
	Glusterfs                     *Io_k8s_api_core_v1_GlusterfsPersistentVolumeSource  `json:"glusterfs,omitempty"`
	HostPath                      *Io_k8s_api_core_v1_HostPathVolumeSource             `json:"hostPath,omitempty"`
	Iscsi                         *Io_k8s_api_core_v1_ISCSIPersistentVolumeSource      `json:"iscsi,omitempty"`
	Local                         *Io_k8s_api_core_v1_LocalVolumeSource                `json:"local,omitempty"`
	MountOptions                  []string                                             `json:"mountOptions,omitempty"`
	Nfs                           *Io_k8s_api_core_v1_NFSVolumeSource                  `json:"nfs,omitempty"`
	NodeAffinity                  *Io_k8s_api_core_v1_VolumeNodeAffinity               `json:"nodeAffinity,omitempty"`
	PersistentVolumeReclaimPolicy string                                               `json:"persistentVolumeReclaimPolicy,omitempty"`
	PhotonPersistentDisk          *Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"`
	PortworxVolume                *Io_k8s_api_core_v1_PortworxVolumeSource             `json:"portworxVolume,omitempty"`
	Quobyte                       *Io_k8s_api_core_v1_QuobyteVolumeSource              `json:"quobyte,omitempty"`
	Rbd                           *Io_k8s_api_core_v1_RBDPersistentVolumeSource        `json:"rbd,omitempty"`
	ScaleIO                       *Io_k8s_api_core_v1_ScaleIOPersistentVolumeSource    `json:"scaleIO,omitempty"`
	StorageClassName              string                                               `json:"storageClassName,omitempty"`
	Storageos                     *Io_k8s_api_core_v1_StorageOSPersistentVolumeSource  `json:"storageos,omitempty"`
	VolumeMode                    string                                               `json:"volumeMode,omitempty"`
	VsphereVolume                 *Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource   `json:"vsphereVolume,omitempty"`
}

