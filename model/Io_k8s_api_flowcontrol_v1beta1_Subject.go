package kugo_model

type Io_k8s_api_flowcontrol_v1beta1_Subject struct {
	Group          *Io_k8s_api_flowcontrol_v1beta1_GroupSubject          `json:"group,omitempty"`
	Kind           string                                                `json:"kind"`
	ServiceAccount *Io_k8s_api_flowcontrol_v1beta1_ServiceAccountSubject `json:"serviceAccount,omitempty"`
	User           *Io_k8s_api_flowcontrol_v1beta1_UserSubject           `json:"user,omitempty"`
}

