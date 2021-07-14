package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type Io_k8s_api_core_v1_AzureFilePersistentVolumeSource struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly        *bool   `json:"readOnly,omitempty"`

	// the name of secret that contains Azure Storage Account Name and Key
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SecretName      *string `json:"secretName"`

	// the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SecretNamespace *string `json:"secretNamespace,omitempty"`

	// Share Name
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ShareName       *string `json:"shareName"`
}
