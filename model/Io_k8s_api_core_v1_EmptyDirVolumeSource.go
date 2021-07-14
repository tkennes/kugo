package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
type Io_k8s_api_core_v1_EmptyDirVolumeSource struct {
	// What type of storage medium should back this directory. The default is "" which means to use the node's default medium.
	// Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Medium    *string                                        `json:"medium,omitempty"`

	// Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium.
	// The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum
	// of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info:
	// http://kubernetes.io/docs/user-guide/volumes#emptydir
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	SizeLimit *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"sizeLimit,omitempty"`
}
