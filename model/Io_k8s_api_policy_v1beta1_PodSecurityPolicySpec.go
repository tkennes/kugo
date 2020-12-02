package kugo_model

type Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec struct {
	AllowPrivilegeEscalation        bool                                                        `json:"allowPrivilegeEscalation,omitempty"`
	AllowedCSIDrivers               []Io_k8s_api_policy_v1beta1_AllowedCSIDriver                `json:"allowedCSIDrivers,omitempty"`
	AllowedCapabilities             []string                                                    `json:"allowedCapabilities,omitempty"`
	AllowedFlexVolumes              []Io_k8s_api_policy_v1beta1_AllowedFlexVolume               `json:"allowedFlexVolumes,omitempty"`
	AllowedHostPaths                []Io_k8s_api_policy_v1beta1_AllowedHostPath                 `json:"allowedHostPaths,omitempty"`
	AllowedProcMountTypes           []string                                                    `json:"allowedProcMountTypes,omitempty"`
	AllowedUnsafeSysctls            []string                                                    `json:"allowedUnsafeSysctls,omitempty"`
	DefaultAddCapabilities          []string                                                    `json:"defaultAddCapabilities,omitempty"`
	DefaultAllowPrivilegeEscalation bool                                                        `json:"defaultAllowPrivilegeEscalation,omitempty"`
	ForbiddenSysctls                []string                                                    `json:"forbiddenSysctls,omitempty"`
	FsGroup                         Io_k8s_api_policy_v1beta1_FSGroupStrategyOptions            `json:"fsGroup"`
	HostIPC                         bool                                                        `json:"hostIPC,omitempty"`
	HostNetwork                     bool                                                        `json:"hostNetwork,omitempty"`
	HostPID                         bool                                                        `json:"hostPID,omitempty"`
	HostPorts                       []Io_k8s_api_policy_v1beta1_HostPortRange                   `json:"hostPorts,omitempty"`
	Privileged                      bool                                                        `json:"privileged,omitempty"`
	ReadOnlyRootFilesystem          bool                                                        `json:"readOnlyRootFilesystem,omitempty"`
	RequiredDropCapabilities        []string                                                    `json:"requiredDropCapabilities,omitempty"`
	RunAsGroup                      *Io_k8s_api_policy_v1beta1_RunAsGroupStrategyOptions        `json:"runAsGroup,omitempty"`
	RunAsUser                       Io_k8s_api_policy_v1beta1_RunAsUserStrategyOptions          `json:"runAsUser"`
	RuntimeClass                    *Io_k8s_api_policy_v1beta1_RuntimeClassStrategyOptions      `json:"runtimeClass,omitempty"`
	SeLinux                         Io_k8s_api_policy_v1beta1_SELinuxStrategyOptions            `json:"seLinux"`
	SupplementalGroups              Io_k8s_api_policy_v1beta1_SupplementalGroupsStrategyOptions `json:"supplementalGroups"`
	Volumes                         []string                                                    `json:"volumes,omitempty"`
}

