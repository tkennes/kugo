package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceList.go


// APIService represents a server for a particular GroupVersion. Name must be "version.group".
type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIService struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                                                    `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                                                    `json:"kind,omitempty"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta                           `json:"metadata,omitempty"`

	// Spec contains information for locating and communicating with a server
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceSpec.go
	Spec       *Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceSpec   `json:"spec,omitempty"`

	// Status contains derived information about an API server
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceStatus.go
	Status     *Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceStatus `json:"status,omitempty"`
}
