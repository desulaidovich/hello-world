package response

import (
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

func Err(w http.ResponseWriter, statusCode int, err *dto.Error) error {
	return write(w, statusCode, &dto.Response{
		Success: false,
		Error:   err,
	})
}
