package response

import (
	"net/http"
)

type ErrorCode string

const (
	BadRequest         ErrorCode = "BAD_REQUEST"
	InvalidRequestBody ErrorCode = "INVALID_REQUEST_BODY"
	InternalError      ErrorCode = "INTERNAL_ERROR"
)

type ErrorInfo struct {
	StatusCode int
	Message    string
}

var ErrorMap = map[ErrorCode]ErrorInfo{
	BadRequest: {
		StatusCode: http.StatusBadRequest,
		Message:    "The request could not be understood due to malformed syntax.",
	},
	InvalidRequestBody: {
		StatusCode: http.StatusBadRequest,
		Message:    "The request body contains invalid JSON or missing required fields.",
	},
	InternalError: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong on our end.",
	},
}
