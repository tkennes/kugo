package kugo_model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EnvVarSource.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_DownwardAPIVolumeFile.go


// ObjectFieldSelector selects an APIVersioned field of an object.
type Io_k8s_api_core_v1_ObjectFieldSelector struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Path of the field to select in the specified API version.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	FieldPath  *string `json:"fieldPath"`
}
