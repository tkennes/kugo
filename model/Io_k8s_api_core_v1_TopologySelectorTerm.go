package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_StorageClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_StorageClass.go


// A topology selector term represents the result of label queries. A null or empty topology selector term matches no
// objects. The requirements of them are ANDed. It provides a subset of functionality as NodeSelectorTerm. This is an alpha
// feature and may change in the future.
type Io_k8s_api_core_v1_TopologySelectorTerm struct {
	// A list of topology selector requirements by labels.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_TopologySelectorLabelRequirement.go
	MatchLabelExpressions []Io_k8s_api_core_v1_TopologySelectorLabelRequirement `json:"matchLabelExpressions,omitempty"`
}
