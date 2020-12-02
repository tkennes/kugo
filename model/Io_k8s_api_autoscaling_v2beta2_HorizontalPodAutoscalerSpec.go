package kugo_model

type Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerSpec struct {
	Behavior       *Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerBehavior `json:"behavior,omitempty"`
	MaxReplicas    int                                                             `json:"maxReplicas"`
	Metrics        []Io_k8s_api_autoscaling_v2beta2_MetricSpec                     `json:"metrics,omitempty"`
	MinReplicas    int                                                             `json:"minReplicas,omitempty"`
	ScaleTargetRef Io_k8s_api_autoscaling_v2beta2_CrossVersionObjectReference      `json:"scaleTargetRef"`
}

