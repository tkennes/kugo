package model

import (
	"time"
)


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1beta1_EventList.go


// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events
// have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the
// timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events
// with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type Io_k8s_api_events_v1beta1_Event struct {
	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at
	// most 128 characters.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Action                   *string                                         `json:"action,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion               *string                                         `json:"apiVersion,omitempty"`

	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DeprecatedCount          *int                                            `json:"deprecatedCount,omitempty"`

	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	DeprecatedFirstTimestamp *time.Time                                      `json:"deprecatedFirstTimestamp,omitempty"`

	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	DeprecatedLastTimestamp  *time.Time                                      `json:"deprecatedLastTimestamp,omitempty"`

	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EventSource.go
	DeprecatedSource         *Io_k8s_api_core_v1_EventSource                 `json:"deprecatedSource,omitempty"`

	// eventTime is the time when this Event was first observed. It is required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	EventTime                *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"eventTime"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind                     *string                                         `json:"kind,omitempty"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata                 Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata"`

	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries
	// should be prepared to handle values up to 64kB.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Note                     *string                                         `json:"note,omitempty"`

	// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason                   *string                                         `json:"reason,omitempty"`

	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g.
	// ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet
	// object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	Regarding                *Io_k8s_api_core_v1_ObjectReference             `json:"regarding,omitempty"`

	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or
	// deletion of related object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	Related                  *Io_k8s_api_core_v1_ObjectReference             `json:"related,omitempty"`

	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field
	// cannot be empty for new Events.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ReportingController      *string                                         `json:"reportingController,omitempty"`

	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events
	// and it can have at most 128 characters.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ReportingInstance        *string                                         `json:"reportingInstance,omitempty"`

	// series is data about the Event series this event represents or nil if it's a singleton Event.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1beta1_EventSeries.go
	Series                   *Io_k8s_api_events_v1beta1_EventSeries          `json:"series,omitempty"`

	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type                     *string                                         `json:"type,omitempty"`
}
