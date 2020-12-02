package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_DeleteOptions struct {
	ApiVersion         string                                              `json:"apiVersion,omitempty"`
	DryRun             []string                                            `json:"dryRun,omitempty"`
	GracePeriodSeconds int                                                 `json:"gracePeriodSeconds,omitempty"`
	Kind               string                                              `json:"kind,omitempty"`
	OrphanDependents   bool                                                `json:"orphanDependents,omitempty"`
	Preconditions      *Io_k8s_apimachinery_pkg_apis_meta_v1_Preconditions `json:"preconditions,omitempty"`
	PropagationPolicy  string                                              `json:"propagationPolicy,omitempty"`
}

