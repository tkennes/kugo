package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set.
// Flocker volumes do not support ownership management or SELinux relabeling.
type Io_k8s_api_core_v1_FlockerVolumeSource struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	DatasetName *string `json:"datasetName,omitempty"`

	// UUID of the dataset. This is unique identifier of a Flocker dataset
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	DatasetUUID *string `json:"datasetUUID,omitempty"`
}
