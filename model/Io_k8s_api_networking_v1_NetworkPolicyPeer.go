package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyEgressRule.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyIngressRule.go


// NetworkPolicyPeer describes a peer to allow traffic to/from. Only certain combinations of fields are allowed
type Io_k8s_api_networking_v1_NetworkPolicyPeer struct {
	// IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IPBlock.go
	IpBlock           *Io_k8s_api_networking_v1_IPBlock                   `json:"ipBlock,omitempty"`

	// Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but
	// empty, it selects all namespaces.  If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods
	// matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces
	// selected by NamespaceSelector.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	NamespaceSelector *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"namespaceSelector,omitempty"`

	// This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty,
	// it selects all pods.  If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching
	// PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the
	// policy's own Namespace.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	PodSelector       *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"podSelector,omitempty"`
}
