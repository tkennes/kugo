package kugo_model

type Io_k8s_api_storage_v1beta1_VolumeAttachmentStatus struct {
	AttachError        *Io_k8s_api_storage_v1beta1_VolumeError `json:"attachError,omitempty"`
	Attached           bool                                    `json:"attached"`
	AttachmentMetadata *interface{}                            `json:"attachmentMetadata,omitempty"`
	DetachError        *Io_k8s_api_storage_v1beta1_VolumeError `json:"detachError,omitempty"`
}

