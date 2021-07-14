package model

import (
	"time"
)


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EventList.go


// Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages
// may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a
// consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as
// informative, best-effort, supplemental data.
type Io_k8s_api_core_v1_Event struct {
	// What action was taken/failed regarding to the Regarding object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Action             *string                                         `json:"action,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion         *string                                         `json:"apiVersion,omitempty"`

	// The number of times this event has occurred.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Count              *int                                            `json:"count,omitempty"`

	// Time when this Event was first observed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	EventTime          *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"eventTime,omitempty"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	FirstTimestamp     *time.Time                                      `json:"firstTimestamp,omitempty"`

	// The object that this event is about.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	InvolvedObject     Io_k8s_api_core_v1_ObjectReference              `json:"involvedObject"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind               *string                                         `json:"kind,omitempty"`

	// The time at which the most recent occurrence of this event was recorded.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastTimestamp      *time.Time                                      `json:"lastTimestamp,omitempty"`

	// A human-readable description of the status of this operation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message            *string                                         `json:"message,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata           Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata"`

	// This should be a short, machine understandable string that gives the reason for the transition into the object's current
	// status.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason             *string                                         `json:"reason,omitempty"`

	// Optional secondary object for more complex actions.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	Related            *Io_k8s_api_core_v1_ObjectReference             `json:"related,omitempty"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ReportingComponent *string                                         `json:"reportingComponent,omitempty"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ReportingInstance  *string                                         `json:"reportingInstance,omitempty"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EventSeries.go
	Series             *Io_k8s_api_core_v1_EventSeries                 `json:"series,omitempty"`

	// The component reporting this event. Should be a short machine understandable string.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EventSource.go
	Source             *Io_k8s_api_core_v1_EventSource                 `json:"source,omitempty"`

	// Type of this event (Normal, Warning), new types could be added in the future
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type               *string                                         `json:"type,omitempty"`
}
