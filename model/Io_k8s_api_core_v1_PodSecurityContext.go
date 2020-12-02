package kugo_model

type Io_k8s_api_core_v1_PodSecurityContext struct {
	FsGroup             int                                               `json:"fsGroup,omitempty"`
	FsGroupChangePolicy string                                            `json:"fsGroupChangePolicy,omitempty"`
	RunAsGroup          int                                               `json:"runAsGroup,omitempty"`
	RunAsNonRoot        bool                                              `json:"runAsNonRoot,omitempty"`
	RunAsUser           int                                               `json:"runAsUser,omitempty"`
	SeLinuxOptions      *Io_k8s_api_core_v1_SELinuxOptions                `json:"seLinuxOptions,omitempty"`
	SeccompProfile      *Io_k8s_api_core_v1_SeccompProfile                `json:"seccompProfile,omitempty"`
	SupplementalGroups  []int                                             `json:"supplementalGroups,omitempty"`
	Sysctls             []Io_k8s_api_core_v1_Sysctl                       `json:"sysctls,omitempty"`
	WindowsOptions      *Io_k8s_api_core_v1_WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}

