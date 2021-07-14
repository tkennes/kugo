package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or
// SELinux relabeling.
type Io_k8s_api_core_v1_GlusterfsPersistentVolumeSource struct {
	// EndpointsName is the endpoint name that details Glusterfs topology. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Endpoints          *string `json:"endpoints"`

	// EndpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace
	// defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-
	// a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	EndpointsNamespace *string `json:"endpointsNamespace,omitempty"`

	// Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path               *string `json:"path"`

	// ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly           *bool   `json:"readOnly,omitempty"`
}
