package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or
// SELinux relabeling.
type Io_k8s_api_core_v1_GlusterfsVolumeSource struct {
	// EndpointsName is the endpoint name that details Glusterfs topology. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Endpoints *string `json:"endpoints"`

	// Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path      *string `json:"path"`

	// ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}
