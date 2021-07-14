package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ProjectedVolumeSource.go


// Projection that may be projected along with other supported volume types
type Io_k8s_api_core_v1_VolumeProjection struct {
	// information about the configMap data to project
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMapProjection.go
	ConfigMap           *Io_k8s_api_core_v1_ConfigMapProjection           `json:"configMap,omitempty"`

	// information about the downwardAPI data to project
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_DownwardAPIProjection.go
	DownwardAPI         *Io_k8s_api_core_v1_DownwardAPIProjection         `json:"downwardAPI,omitempty"`

	// information about the secret data to project
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretProjection.go
	Secret              *Io_k8s_api_core_v1_SecretProjection              `json:"secret,omitempty"`

	// information about the serviceAccountToken data to project
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceAccountTokenProjection.go
	ServiceAccountToken *Io_k8s_api_core_v1_ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"`
}
