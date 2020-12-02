package kugo_model

type Io_k8s_api_core_v1_ServiceSpec struct {
	AllocateLoadBalancerNodePorts bool                                      `json:"allocateLoadBalancerNodePorts,omitempty"`
	ClusterIP                     string                                    `json:"clusterIP,omitempty"`
	ClusterIPs                    []string                                  `json:"clusterIPs,omitempty"`
	ExternalIPs                   []string                                  `json:"externalIPs,omitempty"`
	ExternalName                  string                                    `json:"externalName,omitempty"`
	ExternalTrafficPolicy         string                                    `json:"externalTrafficPolicy,omitempty"`
	HealthCheckNodePort           int                                       `json:"healthCheckNodePort,omitempty"`
	IpFamilies                    []string                                  `json:"ipFamilies,omitempty"`
	IpFamilyPolicy                string                                    `json:"ipFamilyPolicy,omitempty"`
	LoadBalancerIP                string                                    `json:"loadBalancerIP,omitempty"`
	LoadBalancerSourceRanges      []string                                  `json:"loadBalancerSourceRanges,omitempty"`
	Ports                         []Io_k8s_api_core_v1_ServicePort          `json:"ports,omitempty"`
	PublishNotReadyAddresses      bool                                      `json:"publishNotReadyAddresses,omitempty"`
	Selector                      *interface{}                              `json:"selector,omitempty"`
	SessionAffinity               string                                    `json:"sessionAffinity,omitempty"`
	SessionAffinityConfig         *Io_k8s_api_core_v1_SessionAffinityConfig `json:"sessionAffinityConfig,omitempty"`
	TopologyKeys                  []string                                  `json:"topologyKeys,omitempty"`
	Type                          string                                    `json:"type,omitempty"`
}

