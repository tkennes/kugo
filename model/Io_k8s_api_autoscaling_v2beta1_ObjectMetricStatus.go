package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricStatus.go


// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second
// on an Ingress object).
type Io_k8s_api_autoscaling_v2beta1_ObjectMetricStatus struct {
	// averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	AverageValue *Io_k8s_apimachinery_pkg_api_resource_Quantity             `json:"averageValue,omitempty"`

	// currentValue is the current value of the metric (as a quantity).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	CurrentValue *Io_k8s_apimachinery_pkg_api_resource_Quantity             `json:"currentValue"`

	// metricName is the name of the metric in question.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MetricName   *string                                                    `json:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the
	// ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping.
	// When unset, just the metricName will be used to gather metrics.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector     *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector        `json:"selector,omitempty"`

	// target is the described Kubernetes object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_CrossVersionObjectReference.go
	Target       Io_k8s_api_autoscaling_v2beta1_CrossVersionObjectReference `json:"target"`
}
