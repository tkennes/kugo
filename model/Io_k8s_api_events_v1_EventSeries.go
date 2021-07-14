package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1_Event.go


// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How
// often to update the EventSeries is up to the event reporters. The default event reporter in "k8s.io/client-
// go/tools/events/event_broadcaster.go" shows how this struct is updated on heartbeats and can guide customized reporter
// implementations.
type Io_k8s_api_events_v1_EventSeries struct {
	// count is the number of occurrences in this series up to the last heartbeat time.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Count            *int                                            `json:"count"`

	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	LastObservedTime *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"lastObservedTime"`
}
