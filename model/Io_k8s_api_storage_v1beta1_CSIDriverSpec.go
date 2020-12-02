package kugo_model

type Io_k8s_api_storage_v1beta1_CSIDriverSpec struct {
	AttachRequired       bool                                      `json:"attachRequired,omitempty"`
	FsGroupPolicy        string                                    `json:"fsGroupPolicy,omitempty"`
	PodInfoOnMount       bool                                      `json:"podInfoOnMount,omitempty"`
	RequiresRepublish    bool                                      `json:"requiresRepublish,omitempty"`
	StorageCapacity      bool                                      `json:"storageCapacity,omitempty"`
	TokenRequests        []Io_k8s_api_storage_v1beta1_TokenRequest `json:"tokenRequests,omitempty"`
	VolumeLifecycleModes []string                                  `json:"volumeLifecycleModes,omitempty"`
}

