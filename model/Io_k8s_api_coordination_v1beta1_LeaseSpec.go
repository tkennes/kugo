package kugo_model

type Io_k8s_api_coordination_v1beta1_LeaseSpec struct {
	AcquireTime          *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"acquireTime,omitempty"`
	HolderIdentity       string                                          `json:"holderIdentity,omitempty"`
	LeaseDurationSeconds int                                             `json:"leaseDurationSeconds,omitempty"`
	LeaseTransitions     int                                             `json:"leaseTransitions,omitempty"`
	RenewTime            *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"renewTime,omitempty"`
}

