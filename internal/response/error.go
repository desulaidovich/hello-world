package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

func ErrRaw(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("response: encode raw: %w", err)
	}
	return nil
}

func Err(w http.ResponseWriter, statusCode int, err *dto.Error) error {
	return New(statusCode, &dto.Response{
		Success: false,
		Error:   err,
	}).JSON(w)
}
