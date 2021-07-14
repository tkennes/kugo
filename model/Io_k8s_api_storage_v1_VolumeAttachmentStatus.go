package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachment.go


// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type Io_k8s_api_storage_v1_VolumeAttachmentStatus struct {
	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the
	// attach operation, i.e. the external-attacher.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeError.go
	AttachError        *Io_k8s_api_storage_v1_VolumeError `json:"attachError,omitempty"`

	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Attached           *bool                              `json:"attached"`

	// Upon successful attach, this field is populated with any information returned by the attach operation that must be
	// passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	AttachmentMetadata *interface{}                       `json:"attachmentMetadata,omitempty"`

	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the
	// detach operation, i.e. the external-attacher.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeError.go
	DetachError        *Io_k8s_api_storage_v1_VolumeError `json:"detachError,omitempty"`
}
