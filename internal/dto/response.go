package dto

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Message string        `json:"message,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}
