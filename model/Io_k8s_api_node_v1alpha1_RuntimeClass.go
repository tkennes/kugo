package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1alpha1_RuntimeClassList.go


// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which
// container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or
// cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName
// reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type Io_k8s_api_node_v1alpha1_RuntimeClass struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                          `json:"kind,omitempty"`

	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#spec-and-status
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1alpha1_RuntimeClassSpec.go
	Spec       Io_k8s_api_node_v1alpha1_RuntimeClassSpec        `json:"spec"`
}
