package dto

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Meta struct {
	Total      int `json:"total,omitempty"`
	Page       int `json:"page,omitempty"`
	Limit      int `json:"limit,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

type Error struct {
	Code    string        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}
