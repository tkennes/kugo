package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSINodeList.go


// CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode
// object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically
// populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a
// node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version
// is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node
// object.
type Io_k8s_api_storage_v1_CSINode struct {
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

	// metadata.name must be the Kubernetes node name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// spec is the specification of CSINode
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSINodeSpec.go
	Spec       Io_k8s_api_storage_v1_CSINodeSpec                `json:"spec"`
}
