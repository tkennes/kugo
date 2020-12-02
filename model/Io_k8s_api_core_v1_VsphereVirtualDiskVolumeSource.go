package kugo_model

type Io_k8s_api_core_v1_VsphereVirtualDiskVolumeSource struct {
	FsType            string `json:"fsType,omitempty"`
	StoragePolicyID   string `json:"storagePolicyID,omitempty"`
	StoragePolicyName string `json:"storagePolicyName,omitempty"`
	VolumePath        string `json:"volumePath"`
}

