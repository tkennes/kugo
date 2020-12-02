package kugo_model

type Io_k8s_api_storage_v1beta1_CSINodeDriver struct {
	Allocatable  *Io_k8s_api_storage_v1beta1_VolumeNodeResources `json:"allocatable,omitempty"`
	Name         string                                          `json:"name"`
	NodeID       string                                          `json:"nodeID"`
	TopologyKeys []string                                        `json:"topologyKeys,omitempty"`
}

