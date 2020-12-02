package kugo_model

type Io_k8s_api_core_v1_Container struct {
	Args                     []string                                 `json:"args,omitempty"`
	Command                  []string                                 `json:"command,omitempty"`
	Env                      []Io_k8s_api_core_v1_EnvVar              `json:"env,omitempty"`
	EnvFrom                  []Io_k8s_api_core_v1_EnvFromSource       `json:"envFrom,omitempty"`
	Image                    string                                   `json:"image,omitempty"`
	ImagePullPolicy          string                                   `json:"imagePullPolicy,omitempty"`
	Lifecycle                *Io_k8s_api_core_v1_Lifecycle            `json:"lifecycle,omitempty"`
	LivenessProbe            *Io_k8s_api_core_v1_Probe                `json:"livenessProbe,omitempty"`
	Name                     string                                   `json:"name"`
	Ports                    []Io_k8s_api_core_v1_ContainerPort       `json:"ports,omitempty"`
	ReadinessProbe           *Io_k8s_api_core_v1_Probe                `json:"readinessProbe,omitempty"`
	Resources                *Io_k8s_api_core_v1_ResourceRequirements `json:"resources,omitempty"`
	SecurityContext          *Io_k8s_api_core_v1_SecurityContext      `json:"securityContext,omitempty"`
	StartupProbe             *Io_k8s_api_core_v1_Probe                `json:"startupProbe,omitempty"`
	Stdin                    bool                                     `json:"stdin,omitempty"`
	StdinOnce                bool                                     `json:"stdinOnce,omitempty"`
	TerminationMessagePath   string                                   `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy string                                   `json:"terminationMessagePolicy,omitempty"`
	Tty                      bool                                     `json:"tty,omitempty"`
	VolumeDevices            []Io_k8s_api_core_v1_VolumeDevice        `json:"volumeDevices,omitempty"`
	VolumeMounts             []Io_k8s_api_core_v1_VolumeMount         `json:"volumeMounts,omitempty"`
	WorkingDir               string                                   `json:"workingDir,omitempty"`
}

