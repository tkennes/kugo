package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type Io_k8s_api_core_v1_AzureFileVolumeSource struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly   *bool   `json:"readOnly,omitempty"`

	// the name of secret that contains Azure Storage Account Name and Key
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	SecretName *string `json:"secretName"`

	// Share Name
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ShareName  *string `json:"shareName"`
}
