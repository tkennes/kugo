package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMapList.go


// ConfigMap holds configuration data for pods to consume.
type Io_k8s_api_core_v1_ConfigMap struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can
	// contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in
	// the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	BinaryData *interface{}                                     `json:"binaryData,omitempty"`

	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with
	// non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the
	// BinaryData field, this is enforced during validation process.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Data       *interface{}                                     `json:"data,omitempty"`

	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be
	// modified). If not set to true, the field can be modified at any time. Defaulted to nil. This is a beta field enabled by
	// ImmutableEphemeralVolumes feature gate.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Immutable  *bool                                            `json:"immutable,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                          `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
}
