package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIService.go


// APIServiceStatus contains derived information about an API server
type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceStatus struct {
	// Current service state of apiService.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceCondition.go
	Conditions []Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceCondition `json:"conditions,omitempty"`
}
