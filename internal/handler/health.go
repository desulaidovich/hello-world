package handler

import (
	"net/http"
	"time"

	"github.com/desulaidovich/hello-world/internal/dto"
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

func (h *HealthHandler) Say(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		h.log.Info(
			"health check completed",
			zap.Duration("duration", time.Since(start)),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
	}()

	data := map[string]any{
		"status":    http.StatusText(http.StatusOK),
		"timestamp": time.Now().Unix(),
	}

	if err := response.OK(w, data); err != nil {
		h.log.Error(
			"failed to write health response",
			zap.Error(err),
			zap.String("handler", "health.say"),
			zap.String("operation", "response.write"),
		)

		if err = response.Err(w, response.InternalError, dto.ErrorDetail{
			Field:   "health",
			Message: "failed to render health status",
		}); err != nil {
			h.log.Error(
				"critical: failed to write error response",
				zap.Error(err),
				zap.String("handler", "health.say"),
			)
		}
		return
	}
}
