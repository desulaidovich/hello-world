package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/desulaidovich/hello-world/internal/dto"
)

type Response struct {
	status int
	v      *dto.Response
}

func New(status int, v *dto.Response) *Response {
	return &Response{status: status, v: v}
}

func (r *Response) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(r.v); err != nil {
		return fmt.Errorf("response: encode: %w", err)
	}

	w.WriteHeader(r.status)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("response: write: %w", err)
	}

	return nil
}
