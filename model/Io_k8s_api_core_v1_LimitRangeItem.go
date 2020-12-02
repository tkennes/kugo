package kugo_model

type Io_k8s_api_core_v1_LimitRangeItem struct {
	Default              *interface{} `json:"default,omitempty"`
	DefaultRequest       *interface{} `json:"defaultRequest,omitempty"`
	Max                  *interface{} `json:"max,omitempty"`
	MaxLimitRequestRatio *interface{} `json:"maxLimitRequestRatio,omitempty"`
	Min                  *interface{} `json:"min,omitempty"`
	Type                 string       `json:"type"`
}

