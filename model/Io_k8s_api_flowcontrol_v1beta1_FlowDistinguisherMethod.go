package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1beta1_FlowSchemaSpec.go


// FlowDistinguisherMethod specifies the method of a flow distinguisher.
type Io_k8s_api_flowcontrol_v1beta1_FlowDistinguisherMethod struct {
	// `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type *string `json:"type"`
}
