package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1alpha1_VolumeAttachment.go


// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type Io_k8s_api_storage_v1alpha1_VolumeAttachmentSpec struct {
	// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by
	// GetPluginName().
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Attacher *string                                            `json:"attacher"`

	// The node that the volume should be attached to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	NodeName *string                                            `json:"nodeName"`

	// Source represents the volume that should be attached.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1alpha1_VolumeAttachmentSource.go
	Source   Io_k8s_api_storage_v1alpha1_VolumeAttachmentSource `json:"source"`
}
