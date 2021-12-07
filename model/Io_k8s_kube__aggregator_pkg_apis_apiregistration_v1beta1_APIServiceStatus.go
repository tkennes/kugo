package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIService.go


// APIServiceStatus contains derived information about an API server
type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceStatus struct {
	// Current service state of apiService.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceCondition.go
	Conditions []Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceCondition `json:"conditions,omitempty"`
}
