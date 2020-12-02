package kugo_model

type Io_k8s_api_authorization_v1beta1_SubjectAccessReviewSpec struct {
	Extra                 *interface{}                                            `json:"extra,omitempty"`
	Group                 []string                                                `json:"group,omitempty"`
	NonResourceAttributes *Io_k8s_api_authorization_v1beta1_NonResourceAttributes `json:"nonResourceAttributes,omitempty"`
	ResourceAttributes    *Io_k8s_api_authorization_v1beta1_ResourceAttributes    `json:"resourceAttributes,omitempty"`
	Uid                   string                                                  `json:"uid,omitempty"`
	User                  string                                                  `json:"user,omitempty"`
}

