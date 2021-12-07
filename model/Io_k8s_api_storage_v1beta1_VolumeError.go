package kugo_model

import (
	"time"
)


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_VolumeAttachmentStatus.go


// VolumeError captures an error encountered during a volume operation.
type Io_k8s_api_storage_v1beta1_VolumeError struct {
	// String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not
	// contain sensitive information.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message *string    `json:"message,omitempty"`

	// Time the error was encountered.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	Time    *time.Time `json:"time,omitempty"`
}
