package kugo_model

type Io_k8s_api_core_v1_AWSElasticBlockStoreVolumeSource struct {
	FsType    string `json:"fsType,omitempty"`
	Partition int    `json:"partition,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
	VolumeID  string `json:"volumeID"`
}

