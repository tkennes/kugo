package kugo_model

type Io_k8s_api_core_v1_VolumeMount struct {
	MountPath        string `json:"mountPath"`
	MountPropagation string `json:"mountPropagation,omitempty"`
	Name             string `json:"name"`
	ReadOnly         bool   `json:"readOnly,omitempty"`
	SubPath          string `json:"subPath,omitempty"`
	SubPathExpr      string `json:"subPathExpr,omitempty"`
}

