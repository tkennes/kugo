package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfiguration.go


// PriorityLevelConfigurationSpec specifies the configuration of a priority level.
type Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationSpec struct {
	// `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if
	// `type` is `"Limited"`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_LimitedPriorityLevelConfiguration.go
	Limited *Io_k8s_api_flowcontrol_v1alpha1_LimitedPriorityLevelConfiguration `json:"limited,omitempty"`

	// `type` indicates whether this priority level is subject to limitation on request execution.  A value of `"Exempt"` means
	// that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the
	// capacity made available to other priority levels.  A value of `"Limited"` means that (a) requests of this priority level
	// _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority
	// level. Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type    *string                                                            `json:"type"`
}
