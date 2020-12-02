package kugo_model

type Io_k8s_api_core_v1_NodeStatus struct {
	Addresses       []Io_k8s_api_core_v1_NodeAddress        `json:"addresses,omitempty"`
	Allocatable     *interface{}                            `json:"allocatable,omitempty"`
	Capacity        *interface{}                            `json:"capacity,omitempty"`
	Conditions      []Io_k8s_api_core_v1_NodeCondition      `json:"conditions,omitempty"`
	Config          *Io_k8s_api_core_v1_NodeConfigStatus    `json:"config,omitempty"`
	DaemonEndpoints *Io_k8s_api_core_v1_NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`
	Images          []Io_k8s_api_core_v1_ContainerImage     `json:"images,omitempty"`
	NodeInfo        *Io_k8s_api_core_v1_NodeSystemInfo      `json:"nodeInfo,omitempty"`
	Phase           string                                  `json:"phase,omitempty"`
	VolumesAttached []Io_k8s_api_core_v1_AttachedVolume     `json:"volumesAttached,omitempty"`
	VolumesInUse    []string                                `json:"volumesInUse,omitempty"`
}

