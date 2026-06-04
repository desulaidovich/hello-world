package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

func OKRaw(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("response: encode raw: %w", err)
	}
	return nil
}

func OK(w http.ResponseWriter, data any) error {
	return New(http.StatusOK, &dto.Response{
		Success: true,
		Data:    data,
	}).JSON(w)
}
