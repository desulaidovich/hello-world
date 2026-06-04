package handler

import (
	"net/http"

	"github.com/desulaidovich/hello-world/internal/response"
	"go.uber.org/zap"
)

type HealthHandler struct {
	log *zap.Logger
}

func NewHealth(log *zap.Logger) *HealthHandler {
	return &HealthHandler{
		log: log,
	}
}

func (h *HealthHandler) Say(w http.ResponseWriter, _ *http.Request) {
	body := map[string]any{
		"status": http.StatusText(http.StatusOK),
	}

	if err := response.OKRaw(w, body); err != nil {
		h.log.Error(
			"write liveness response",
			zap.Error(err),
		)
	}
}
