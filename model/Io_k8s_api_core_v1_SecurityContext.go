package kugo_model

type Io_k8s_api_core_v1_SecurityContext struct {
	AllowPrivilegeEscalation bool                                              `json:"allowPrivilegeEscalation,omitempty"`
	Capabilities             *Io_k8s_api_core_v1_Capabilities                  `json:"capabilities,omitempty"`
	Privileged               bool                                              `json:"privileged,omitempty"`
	ProcMount                string                                            `json:"procMount,omitempty"`
	ReadOnlyRootFilesystem   bool                                              `json:"readOnlyRootFilesystem,omitempty"`
	RunAsGroup               int                                               `json:"runAsGroup,omitempty"`
	RunAsNonRoot             bool                                              `json:"runAsNonRoot,omitempty"`
	RunAsUser                int                                               `json:"runAsUser,omitempty"`
	SeLinuxOptions           *Io_k8s_api_core_v1_SELinuxOptions                `json:"seLinuxOptions,omitempty"`
	SeccompProfile           *Io_k8s_api_core_v1_SeccompProfile                `json:"seccompProfile,omitempty"`
	WindowsOptions           *Io_k8s_api_core_v1_WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}

