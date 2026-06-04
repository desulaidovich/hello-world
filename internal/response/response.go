package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

func write(w http.ResponseWriter, status int, v *dto.Response) error {
	w.Header().Set("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return fmt.Errorf("response: encode: %w", err)
	}

	w.WriteHeader(status)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("response: write: %w", err)
	}

	return nil
}
