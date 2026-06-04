package response

import (
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

func Err(w http.ResponseWriter, errorCode ErrorCode, details ...dto.ErrorDetail) error {
	code, ok := ErrorMap[errorCode]
	if !ok {
		code = ErrorMap[InternalError]
	}

	err := &dto.Error{
		Message: code.Message,
	}

	if len(details) > 0 {
		err.Details = details
	}

	return write(w, code.StatusCode, &dto.Response{
		Success: false,
		Error:   err,
	})
}
