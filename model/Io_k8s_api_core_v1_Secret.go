package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretList.go


// Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than
// MaxSecretSize bytes.
type Io_k8s_api_core_v1_Secret struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of
	// the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described
	// in https://tools.ietf.org/html/rfc4648#section-4
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Data       *interface{}                                     `json:"data,omitempty"`

	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be
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

	// stringData allows specifying non-binary secret data in string form. It is provided as a write-only convenience method.
	// All keys and values are merged into the data field on write, overwriting any existing values. It is never output when
	// reading from the API.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	StringData *interface{}                                     `json:"stringData,omitempty"`

	// Used to facilitate programmatic handling of secret data.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type       *string                                          `json:"type,omitempty"`
}
