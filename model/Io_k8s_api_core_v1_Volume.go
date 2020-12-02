package kugo_model

type Io_k8s_api_core_v1_Volume struct {
	AwsElasticBlockStore  *Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource  `json:"awsElasticBlockStore,omitempty"`
	AzureDisk             *Io_k8s_api_core_v1_AzureDiskVolumeSource             `json:"azureDisk,omitempty"`
	AzureFile             *Io_k8s_api_core_v1_AzureFileVolumeSource             `json:"azureFile,omitempty"`
	Cephfs                *Io_k8s_api_core_v1_CephFSVolumeSource                `json:"cephfs,omitempty"`
	Cinder                *Io_k8s_api_core_v1_CinderVolumeSource                `json:"cinder,omitempty"`
	ConfigMap             *Io_k8s_api_core_v1_ConfigMapVolumeSource             `json:"configMap,omitempty"`
	Csi                   *Io_k8s_api_core_v1_CSIVolumeSource                   `json:"csi,omitempty"`
	DownwardAPI           *Io_k8s_api_core_v1_DownwardAPIVolumeSource           `json:"downwardAPI,omitempty"`
	EmptyDir              *Io_k8s_api_core_v1_EmptyDirVolumeSource              `json:"emptyDir,omitempty"`
	Ephemeral             *Io_k8s_api_core_v1_EphemeralVolumeSource             `json:"ephemeral,omitempty"`
	Fc                    *Io_k8s_api_core_v1_FCVolumeSource                    `json:"fc,omitempty"`
	FlexVolume            *Io_k8s_api_core_v1_FlexVolumeSource                  `json:"flexVolume,omitempty"`
	Flocker               *Io_k8s_api_core_v1_FlockerVolumeSource               `json:"flocker,omitempty"`
	GcePersistentDisk     *Io_k8s_api_core_v1_GCEPersistentDiskVolumeSource     `json:"gcePersistentDisk,omitempty"`
	GitRepo               *Io_k8s_api_core_v1_GitRepoVolumeSource               `json:"gitRepo,omitempty"`
	Glusterfs             *Io_k8s_api_core_v1_GlusterfsVolumeSource             `json:"glusterfs,omitempty"`
	HostPath              *Io_k8s_api_core_v1_HostPathVolumeSource              `json:"hostPath,omitempty"`
	Iscsi                 *Io_k8s_api_core_v1_ISCSIVolumeSource                 `json:"iscsi,omitempty"`
	Name                  string                                                `json:"name"`
	Nfs                   *Io_k8s_api_core_v1_NFSVolumeSource                   `json:"nfs,omitempty"`
	PersistentVolumeClaim *Io_k8s_api_core_v1_PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
	PhotonPersistentDisk  *Io_k8s_api_core_v1_PhotonPersistentDiskVolumeSource  `json:"photonPersistentDisk,omitempty"`
	PortworxVolume        *Io_k8s_api_core_v1_PortworxVolumeSource              `json:"portworxVolume,omitempty"`
	Projected             *Io_k8s_api_core_v1_ProjectedVolumeSource             `json:"projected,omitempty"`
	Quobyte               *Io_k8s_api_core_v1_QuobyteVolumeSource               `json:"quobyte,omitempty"`
	Rbd                   *Io_k8s_api_core_v1_RBDVolumeSource                   `json:"rbd,omitempty"`
	ScaleIO               *Io_k8s_api_core_v1_ScaleIOVolumeSource               `json:"scaleIO,omitempty"`
	Secret                *Io_k8s_api_core_v1_SecretVolumeSource                `json:"secret,omitempty"`
	Storageos             *Io_k8s_api_core_v1_StorageOSVolumeSource             `json:"storageos,omitempty"`
	VsphereVolume         *Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource    `json:"vsphereVolume,omitempty"`
}

