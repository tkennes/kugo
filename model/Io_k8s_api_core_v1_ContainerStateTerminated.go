package kugo_model

import (
	"time"
)

type Io_k8s_api_core_v1_ContainerStateTerminated struct {
	ContainerID string     `json:"containerID,omitempty"`
	ExitCode    int        `json:"exitCode"`
	FinishedAt  *time.Time `json:"finishedAt,omitempty"`
	Message     string     `json:"message,omitempty"`
	Reason      string     `json:"reason,omitempty"`
	Signal      int        `json:"signal,omitempty"`
	StartedAt   *time.Time `json:"startedAt,omitempty"`
}

