package kugo_model

type Io_k8s_api_policy_v1beta1_PodDisruptionBudgetStatus struct {
	CurrentHealthy     int          `json:"currentHealthy"`
	DesiredHealthy     int          `json:"desiredHealthy"`
	DisruptedPods      *interface{} `json:"disruptedPods,omitempty"`
	DisruptionsAllowed int          `json:"disruptionsAllowed"`
	ExpectedPods       int          `json:"expectedPods"`
	ObservedGeneration int          `json:"observedGeneration,omitempty"`
}

