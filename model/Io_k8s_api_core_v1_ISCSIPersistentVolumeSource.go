package kugo_model

type Io_k8s_api_core_v1_ISCSIPersistentVolumeSource struct {
	ChapAuthDiscovery bool                                `json:"chapAuthDiscovery,omitempty"`
	ChapAuthSession   bool                                `json:"chapAuthSession,omitempty"`
	FsType            string                              `json:"fsType,omitempty"`
	InitiatorName     string                              `json:"initiatorName,omitempty"`
	Iqn               string                              `json:"iqn"`
	IscsiInterface    string                              `json:"iscsiInterface,omitempty"`
	Lun               int                                 `json:"lun"`
	Portals           []string                            `json:"portals,omitempty"`
	ReadOnly          bool                                `json:"readOnly,omitempty"`
	SecretRef         *Io_k8s_api_core_v1_SecretReference `json:"secretRef,omitempty"`
	TargetPortal      string                              `json:"targetPortal"`
}

