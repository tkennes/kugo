package kugo_model

type Io_k8s_api_storage_v1beta1_VolumeAttachmentSpec struct {
	Attacher string                                            `json:"attacher"`
	NodeName string                                            `json:"nodeName"`
	Source   Io_k8s_api_storage_v1beta1_VolumeAttachmentSource `json:"source"`
}

