package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Event.go


// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type Io_k8s_api_core_v1_EventSeries struct {
	// Number of occurrences in this series up to the last heartbeat time
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Count            *int                                            `json:"count,omitempty"`

	// Time of the last occurrence observed
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	LastObservedTime *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"lastObservedTime,omitempty"`
}
