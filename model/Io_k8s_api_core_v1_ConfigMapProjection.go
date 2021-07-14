package model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VolumeProjection.go


// Adapts a ConfigMap into a projected volume.  The contents of the target ConfigMap's Data field will be presented in a
// projected volume as files using the keys in the Data field as the file names, unless the items element is populated with
// specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.
type Io_k8s_api_core_v1_ConfigMapProjection struct {
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a
	// file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified
	// paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume
	// setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with
	// '..'.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_KeyToPath.go
	Items    []Io_k8s_api_core_v1_KeyToPath `json:"items,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name     *string                        `json:"name,omitempty"`

	// Specify whether the ConfigMap or its keys must be defined
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Optional *bool                          `json:"optional,omitempty"`
}
