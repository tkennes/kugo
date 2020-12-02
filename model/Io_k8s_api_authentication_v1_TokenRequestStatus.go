package kugo_model

import (
	"time"
)

type Io_k8s_api_authentication_v1_TokenRequestStatus struct {
	ExpirationTimestamp time.Time `json:"expirationTimestamp"`
	Token               string    `json:"token"`
}

