package kugo_model

type Io_k8s_api_core_v1_ServiceStatus struct {
	Conditions   []Io_k8s_apimachinery_pkg_apis_meta_v1_Condition `json:"conditions,omitempty"`
	LoadBalancer *Io_k8s_api_core_v1_LoadBalancerStatus           `json:"loadBalancer,omitempty"`
}

