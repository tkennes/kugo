package kugo_model

import (
	"time"
)

type Io_k8s_api_events_v1beta1_Event struct {
	Action                   string                                          `json:"action,omitempty"`
	ApiVersion               string                                          `json:"apiVersion,omitempty"`
	DeprecatedCount          int                                             `json:"deprecatedCount,omitempty"`
	DeprecatedFirstTimestamp *time.Time                                      `json:"deprecatedFirstTimestamp,omitempty"`
	DeprecatedLastTimestamp  *time.Time                                      `json:"deprecatedLastTimestamp,omitempty"`
	DeprecatedSource         *Io_k8s_api_core_v1_EventSource                 `json:"deprecatedSource,omitempty"`
	EventTime                Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime  `json:"eventTime"`
	Kind                     string                                          `json:"kind,omitempty"`
	Metadata                 Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata"`
	Note                     string                                          `json:"note,omitempty"`
	Reason                   string                                          `json:"reason,omitempty"`
	Regarding                *Io_k8s_api_core_v1_ObjectReference             `json:"regarding,omitempty"`
	Related                  *Io_k8s_api_core_v1_ObjectReference             `json:"related,omitempty"`
	ReportingController      string                                          `json:"reportingController,omitempty"`
	ReportingInstance        string                                          `json:"reportingInstance,omitempty"`
	Series                   *Io_k8s_api_events_v1beta1_EventSeries          `json:"series,omitempty"`
	Type                     string                                          `json:"type,omitempty"`
}

