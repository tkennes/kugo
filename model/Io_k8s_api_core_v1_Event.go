package kugo_model

import (
	"time"
)

type Io_k8s_api_core_v1_Event struct {
	Action             string                                          `json:"action,omitempty"`
	ApiVersion         string                                          `json:"apiVersion,omitempty"`
	Count              int                                             `json:"count,omitempty"`
	EventTime          *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"eventTime,omitempty"`
	FirstTimestamp     *time.Time                                      `json:"firstTimestamp,omitempty"`
	InvolvedObject     Io_k8s_api_core_v1_ObjectReference              `json:"involvedObject"`
	Kind               string                                          `json:"kind,omitempty"`
	LastTimestamp      *time.Time                                      `json:"lastTimestamp,omitempty"`
	Message            string                                          `json:"message,omitempty"`
	Metadata           Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata"`
	Reason             string                                          `json:"reason,omitempty"`
	Related            *Io_k8s_api_core_v1_ObjectReference             `json:"related,omitempty"`
	ReportingComponent string                                          `json:"reportingComponent,omitempty"`
	ReportingInstance  string                                          `json:"reportingInstance,omitempty"`
	Series             *Io_k8s_api_core_v1_EventSeries                 `json:"series,omitempty"`
	Source             *Io_k8s_api_core_v1_EventSource                 `json:"source,omitempty"`
	Type               string                                          `json:"type,omitempty"`
}

