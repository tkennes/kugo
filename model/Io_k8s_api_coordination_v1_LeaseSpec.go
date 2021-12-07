package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_coordination_v1_Lease.go


// LeaseSpec is a specification of a Lease.
type Io_k8s_api_coordination_v1_LeaseSpec struct {
	// acquireTime is a time when the current lease was acquired.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	AcquireTime          *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"acquireTime,omitempty"`

	// holderIdentity contains the identity of the holder of a current lease.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	HolderIdentity       *string                                         `json:"holderIdentity,omitempty"`

	// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against
	// time of last observed RenewTime.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	LeaseDurationSeconds *int                                            `json:"leaseDurationSeconds,omitempty"`

	// leaseTransitions is the number of transitions of a lease between holders.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	LeaseTransitions     *int                                            `json:"leaseTransitions,omitempty"`

	// renewTime is a time when the current holder of a lease has last updated the lease.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime.go
	RenewTime            *Io_k8s_apimachinery_pkg_apis_meta_v1_MicroTime `json:"renewTime,omitempty"`
}
