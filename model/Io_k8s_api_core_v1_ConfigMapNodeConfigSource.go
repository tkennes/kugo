package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeConfigSource.go


// ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node.
type Io_k8s_api_core_v1_ConfigMapNodeConfigSource struct {
	// KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This
	// field is required in all cases.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	KubeletConfigKey *string `json:"kubeletConfigKey"`

	// Name is the metadata.name of the referenced ConfigMap. This field is required in all cases.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name             *string `json:"name"`

	// Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Namespace        *string `json:"namespace"`

	// ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and
	// required in Node.Status.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ResourceVersion  *string `json:"resourceVersion,omitempty"`

	// UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Uid              *string `json:"uid,omitempty"`
}
