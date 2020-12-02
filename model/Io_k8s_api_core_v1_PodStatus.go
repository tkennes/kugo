package kugo_model

import (
	"time"
)

type Io_k8s_api_core_v1_PodStatus struct {
	Conditions                 []Io_k8s_api_core_v1_PodCondition    `json:"conditions,omitempty"`
	ContainerStatuses          []Io_k8s_api_core_v1_ContainerStatus `json:"containerStatuses,omitempty"`
	EphemeralContainerStatuses []Io_k8s_api_core_v1_ContainerStatus `json:"ephemeralContainerStatuses,omitempty"`
	HostIP                     string                               `json:"hostIP,omitempty"`
	InitContainerStatuses      []Io_k8s_api_core_v1_ContainerStatus `json:"initContainerStatuses,omitempty"`
	Message                    string                               `json:"message,omitempty"`
	NominatedNodeName          string                               `json:"nominatedNodeName,omitempty"`
	Phase                      string                               `json:"phase,omitempty"`
	PodIP                      string                               `json:"podIP,omitempty"`
	PodIPs                     []Io_k8s_api_core_v1_PodIP           `json:"podIPs,omitempty"`
	QosClass                   string                               `json:"qosClass,omitempty"`
	Reason                     string                               `json:"reason,omitempty"`
	StartTime                  *time.Time                           `json:"startTime,omitempty"`
}

