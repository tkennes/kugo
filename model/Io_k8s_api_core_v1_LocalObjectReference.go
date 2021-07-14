package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ISCSIVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ScaleIOVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CSIVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CinderVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_RBDVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceAccount.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_StorageOSVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_FlexVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_CephFSVolumeSource.go


// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type Io_k8s_api_core_v1_LocalObjectReference struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name *string `json:"name,omitempty"`
}
