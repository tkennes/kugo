package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_Subject.go


// ServiceAccountSubject holds detailed information for service-account-kind subject.
type Io_k8s_api_flowcontrol_v1beta1_ServiceAccountSubject struct {
	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name      *string `json:"name"`

	// `namespace` is the namespace of matching ServiceAccount objects. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace *string `json:"namespace"`
}
