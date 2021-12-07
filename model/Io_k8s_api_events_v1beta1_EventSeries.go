package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_events_v1beta1_Event.go


// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type Io_k8s_api_events_v1beta1_EventSeries struct {
	// count is the number of occurrences in this series up to the last heartbeat time.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Count            *int                                            `json:"count"`

	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	LastObservedTime *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"lastObservedTime"`
}
