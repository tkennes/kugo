package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicy.go


// NetworkPolicySpec provides the specification of a NetworkPolicy
type Io_k8s_api_networking_v1_NetworkPolicySpec struct {
	// List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies
	// selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule
	// across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this
	// NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by
	// default). This field is beta-level in 1.8
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyEgressRule.go
	Egress      []Io_k8s_api_networking_v1_NetworkPolicyEgressRule  `json:"egress,omitempty"`

	// List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies
	// selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node,
	// OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches
	// the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the
	// pods it selects are isolated by default)
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyIngressRule.go
	Ingress     []Io_k8s_api_networking_v1_NetworkPolicyIngressRule `json:"ingress,omitempty"`

	// Selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected
	// by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are
	// combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector
	// matches all pods in this namespace.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	PodSelector Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector  `json:"podSelector"`

	// List of rule types that the NetworkPolicy relates to. Valid options are "Ingress", "Egress", or "Ingress,Egress". If
	// this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an
	// Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are
	// assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [
	// "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a
	// policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise
	// default to just [ "Ingress" ]). This field is beta-level in 1.8
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PolicyTypes []*string                                           `json:"policyTypes,omitempty"`
}
