package kugo_model

type Io_k8s_api_autoscaling_v1_HorizontalPodAutoscalerSpec struct {
	MaxReplicas                    int                                                   `json:"maxReplicas"`
	MinReplicas                    int                                                   `json:"minReplicas,omitempty"`
	ScaleTargetRef                 Io_k8s_api_autoscaling_v1_CrossVersionObjectReference `json:"scaleTargetRef"`
	TargetCPUUtilizationPercentage int                                                   `json:"targetCPUUtilizationPercentage,omitempty"`
}

