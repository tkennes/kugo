package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_HPAScalingRules.go


// HPAScalingPolicy is a single policy which must hold true for a specified past interval.
type Io_k8s_api_autoscaling_v2beta2_HPAScalingPolicy struct {
	// PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than
	// zero and less than or equal to 1800 (30 min).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	PeriodSeconds *int    `json:"periodSeconds"`

	// Type is used to specify the scaling policy.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type          *string `json:"type"`

	// Value contains the amount of change which is permitted by the policy. It must be greater than zero
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Value         *int    `json:"value"`
}
