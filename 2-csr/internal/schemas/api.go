package schemas

type APIResponse[K any] struct {
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      K      `json:"data,omitempty"`
}
