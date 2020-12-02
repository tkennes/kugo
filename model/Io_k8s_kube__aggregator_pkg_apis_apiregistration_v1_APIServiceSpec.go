package kugo_model

type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceSpec struct {
	CaBundle              string                                                                `json:"caBundle,omitempty"`
	Group                 string                                                                `json:"group,omitempty"`
	GroupPriorityMinimum  int                                                                   `json:"groupPriorityMinimum"`
	InsecureSkipTLSVerify bool                                                                  `json:"insecureSkipTLSVerify,omitempty"`
	Service               *Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_ServiceReference `json:"service,omitempty"`
	Version               string                                                                `json:"version,omitempty"`
	VersionPriority       int                                                                   `json:"versionPriority"`
}

