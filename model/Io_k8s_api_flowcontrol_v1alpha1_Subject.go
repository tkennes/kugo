package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PolicyRulesWithSubjects.go


// Subject matches the originator of a request, as identified by the request authentication system. There are three ways of
// matching an originator; by user, group, or service account.
type Io_k8s_api_flowcontrol_v1alpha1_Subject struct {
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_GroupSubject.go
	Group          *Io_k8s_api_flowcontrol_v1alpha1_GroupSubject          `json:"group,omitempty"`

	// Required
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind           *string                                                `json:"kind"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_ServiceAccountSubject.go
	ServiceAccount *Io_k8s_api_flowcontrol_v1alpha1_ServiceAccountSubject `json:"serviceAccount,omitempty"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_UserSubject.go
	User           *Io_k8s_api_flowcontrol_v1alpha1_UserSubject           `json:"user,omitempty"`
}
