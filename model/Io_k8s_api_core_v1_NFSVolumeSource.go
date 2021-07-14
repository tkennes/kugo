package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux
// relabeling.
type Io_k8s_api_core_v1_NFSVolumeSource struct {
	// Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path     *string `json:"path"`

	// ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly *bool   `json:"readOnly,omitempty"`

	// Server is the hostname or IP address of the NFS server. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Server   *string `json:"server"`
}
