package schemas

import "time"

type APIResponse[K any] struct {
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      K      `json:"data,omitempty"`
}

func NewAPIMessageResponse(message string) APIResponse[any] {
	return APIResponse[any]{
		Message:   message,
		Timestamp: time.Now().Unix(),
	}
}

func NewAPIResponse[K any](message string, data K) APIResponse[K] {
	return APIResponse[K]{
		Message:   message,
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
}
