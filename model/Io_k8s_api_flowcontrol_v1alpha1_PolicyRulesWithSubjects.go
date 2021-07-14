package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_FlowSchemaSpec.go


// PolicyRulesWithSubjects prescribes a test that applies to a request to an apiserver. The test considers the subject
// making the request, the verb being requested, and the resource to be acted upon. This PolicyRulesWithSubjects matches a
// request if and only if both (a) at least one member of subjects matches the request and (b) at least one member of
// resourceRules or nonResourceRules matches the request.
type Io_k8s_api_flowcontrol_v1alpha1_PolicyRulesWithSubjects struct {
	// `nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb and the
	// target non-resource URL.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_NonResourcePolicyRule.go
	NonResourceRules []Io_k8s_api_flowcontrol_v1alpha1_NonResourcePolicyRule `json:"nonResourceRules,omitempty"`

	// `resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target
	// resource. At least one of `resourceRules` and `nonResourceRules` has to be non-empty.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_ResourcePolicyRule.go
	ResourceRules    []Io_k8s_api_flowcontrol_v1alpha1_ResourcePolicyRule    `json:"resourceRules,omitempty"`

	// subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one
	// member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches
	// every request. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_Subject.go
	Subjects         []Io_k8s_api_flowcontrol_v1alpha1_Subject               `json:"subjects"`
}
