package kugo_model

type Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerSpec struct {
	MaxReplicas    int                                                        `json:"maxReplicas"`
	Metrics        []Io_k8s_api_autoscaling_v2beta1_MetricSpec                `json:"metrics,omitempty"`
	MinReplicas    int                                                        `json:"minReplicas,omitempty"`
	ScaleTargetRef Io_k8s_api_autoscaling_v2beta1_CrossVersionObjectReference `json:"scaleTargetRef"`
}

