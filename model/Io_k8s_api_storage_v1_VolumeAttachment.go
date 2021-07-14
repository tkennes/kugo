package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachmentList.go


// VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.
// VolumeAttachment objects are non-namespaced.
type Io_k8s_api_storage_v1_VolumeAttachment struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                          `json:"kind,omitempty"`

	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachmentSpec.go
	Spec       Io_k8s_api_storage_v1_VolumeAttachmentSpec       `json:"spec"`

	// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the
	// external-attacher.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachmentStatus.go
	Status     *Io_k8s_api_storage_v1_VolumeAttachmentStatus    `json:"status,omitempty"`
}
