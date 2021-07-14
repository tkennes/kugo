package kugo_model

import (
	"time"
)


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerState.go


// ContainerStateTerminated is a terminated state of a container.
type Io_k8s_api_core_v1_ContainerStateTerminated struct {
	// Container's ID in the format 'docker://<container_id>'
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ContainerID *string    `json:"containerID,omitempty"`

	// Exit status from the last termination of the container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ExitCode    *int       `json:"exitCode"`

	// Time at which the container last terminated
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	FinishedAt  *time.Time `json:"finishedAt,omitempty"`

	// Message regarding the last termination of the container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message     *string    `json:"message,omitempty"`

	// (brief) reason from the last termination of the container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason      *string    `json:"reason,omitempty"`

	// Signal from the last termination of the container
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Signal      *int       `json:"signal,omitempty"`

	// Time at which previous execution of the container started
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	StartedAt   *time.Time `json:"startedAt,omitempty"`
}
