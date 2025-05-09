package global

type Response[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
	Error   *any   `json:"error,omitempty"`
}
