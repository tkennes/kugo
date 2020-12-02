package kugo_model

import (
	"time"
)

type Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta struct {
	Annotations                *interface{}                                              `json:"annotations,omitempty"`
	ClusterName                string                                                    `json:"clusterName,omitempty"`
	CreationTimestamp          *time.Time                                                `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds int                                                       `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp          *time.Time                                                `json:"deletionTimestamp,omitempty"`
	Finalizers                 []string                                                  `json:"finalizers,omitempty"`
	GenerateName               string                                                    `json:"generateName,omitempty"`
	Generation                 int                                                       `json:"generation,omitempty"`
	Labels                     *interface{}                                              `json:"labels,omitempty"`
	ManagedFields              []Io_k8s_apimachinery_pkg_apis_meta_v1_ManagedFieldsEntry `json:"managedFields,omitempty"`
	Name                       string                                                    `json:"name,omitempty"`
	Namespace                  string                                                    `json:"namespace,omitempty"`
	OwnerReferences            []Io_k8s_apimachinery_pkg_apis_meta_v1_OwnerReference     `json:"ownerReferences,omitempty"`
	ResourceVersion            string                                                    `json:"resourceVersion,omitempty"`
	SelfLink                   string                                                    `json:"selfLink,omitempty"`
	Uid                        string                                                    `json:"uid,omitempty"`
}

