package kugo_model

type Io_k8s_api_core_v1_PodSpec struct {
	ActiveDeadlineSeconds         int                                           `json:"activeDeadlineSeconds,omitempty"`
	Affinity                      *Io_k8s_api_core_v1_Affinity                  `json:"affinity,omitempty"`
	AutomountServiceAccountToken  bool                                          `json:"automountServiceAccountToken,omitempty"`
	Containers                    []Io_k8s_api_core_v1_Container                `json:"containers"`
	DnsConfig                     *Io_k8s_api_core_v1_PodDNSConfig              `json:"dnsConfig,omitempty"`
	DnsPolicy                     string                                        `json:"dnsPolicy,omitempty"`
	EnableServiceLinks            bool                                          `json:"enableServiceLinks,omitempty"`
	EphemeralContainers           []Io_k8s_api_core_v1_EphemeralContainer       `json:"ephemeralContainers,omitempty"`
	HostAliases                   []Io_k8s_api_core_v1_HostAlias                `json:"hostAliases,omitempty"`
	HostIPC                       bool                                          `json:"hostIPC,omitempty"`
	HostNetwork                   bool                                          `json:"hostNetwork,omitempty"`
	HostPID                       bool                                          `json:"hostPID,omitempty"`
	Hostname                      string                                        `json:"hostname,omitempty"`
	ImagePullSecrets              []Io_k8s_api_core_v1_LocalObjectReference     `json:"imagePullSecrets,omitempty"`
	InitContainers                []Io_k8s_api_core_v1_Container                `json:"initContainers,omitempty"`
	NodeName                      string                                        `json:"nodeName,omitempty"`
	NodeSelector                  *interface{}                                  `json:"nodeSelector,omitempty"`
	Overhead                      *interface{}                                  `json:"overhead,omitempty"`
	PreemptionPolicy              string                                        `json:"preemptionPolicy,omitempty"`
	Priority                      int                                           `json:"priority,omitempty"`
	PriorityClassName             string                                        `json:"priorityClassName,omitempty"`
	ReadinessGates                []Io_k8s_api_core_v1_PodReadinessGate         `json:"readinessGates,omitempty"`
	RestartPolicy                 string                                        `json:"restartPolicy,omitempty"`
	RuntimeClassName              string                                        `json:"runtimeClassName,omitempty"`
	SchedulerName                 string                                        `json:"schedulerName,omitempty"`
	SecurityContext               *Io_k8s_api_core_v1_PodSecurityContext        `json:"securityContext,omitempty"`
	ServiceAccount                string                                        `json:"serviceAccount,omitempty"`
	ServiceAccountName            string                                        `json:"serviceAccountName,omitempty"`
	SetHostnameAsFQDN             bool                                          `json:"setHostnameAsFQDN,omitempty"`
	ShareProcessNamespace         bool                                          `json:"shareProcessNamespace,omitempty"`
	Subdomain                     string                                        `json:"subdomain,omitempty"`
	TerminationGracePeriodSeconds int                                           `json:"terminationGracePeriodSeconds,omitempty"`
	Tolerations                   []Io_k8s_api_core_v1_Toleration               `json:"tolerations,omitempty"`
	TopologySpreadConstraints     []Io_k8s_api_core_v1_TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
	Volumes                       []Io_k8s_api_core_v1_Volume                   `json:"volumes,omitempty"`
}

