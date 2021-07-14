package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_FlexPersistentVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ISCSIPersistentVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CinderPersistentVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ScaleIOPersistentVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CSIPersistentVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CephFSPersistentVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_RBDPersistentVolumeSource.go


// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
type Io_k8s_api_core_v1_SecretReference struct {
	// Name is unique within a namespace to reference a secret resource.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name      *string `json:"name,omitempty"`

	// Namespace defines the space within which the secret name must be unique.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace *string `json:"namespace,omitempty"`
}
