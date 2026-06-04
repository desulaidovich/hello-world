package response

import (
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

func OK(w http.ResponseWriter, data any) error {
	return write(w, http.StatusOK, &dto.Response{
		Success: true,
		Data:    data,
	})
}
