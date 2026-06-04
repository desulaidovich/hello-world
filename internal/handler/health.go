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
	if err := response.OK(w, map[string]any{
		"status": http.StatusText(http.StatusOK),
	}); err != nil {
		h.log.Error("write liveness response", zap.Error(err))
	}
}
