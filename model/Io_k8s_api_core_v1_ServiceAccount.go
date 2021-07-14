package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceAccountList.go


// ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a
// principal that can be authenticated and authorized * a set of secrets
type Io_k8s_api_core_v1_ServiceAccount struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion                   *string                                          `json:"apiVersion,omitempty"`

	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token
	// automatically mounted. Can be overridden at the pod level.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	AutomountServiceAccountToken *bool                                            `json:"automountServiceAccountToken,omitempty"`

	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that
	// reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but
	// ImagePullSecrets are only accessed by the kubelet. More info:
	// https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LocalObjectReference.go
	ImagePullSecrets             []Io_k8s_api_core_v1_LocalObjectReference        `json:"imagePullSecrets,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind                         *string                                          `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata                     *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Secrets is the list of secrets allowed to be used by pods running using this ServiceAccount. More info:
	// https://kubernetes.io/docs/concepts/configuration/secret
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	Secrets                      []Io_k8s_api_core_v1_ObjectReference             `json:"secrets,omitempty"`
}
