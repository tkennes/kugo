package kugo_model

type Io_k8s_api_authorization_v1_SubjectAccessReviewSpec struct {
	Extra                 *interface{}                                       `json:"extra,omitempty"`
	Groups                []string                                           `json:"groups,omitempty"`
	NonResourceAttributes *Io_k8s_api_authorization_v1_NonResourceAttributes `json:"nonResourceAttributes,omitempty"`
	ResourceAttributes    *Io_k8s_api_authorization_v1_ResourceAttributes    `json:"resourceAttributes,omitempty"`
	Uid                   string                                             `json:"uid,omitempty"`
	User                  string                                             `json:"user,omitempty"`
}

