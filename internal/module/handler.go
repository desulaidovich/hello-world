package module

import (
	"net/http"

	"github.com/desulaidovich/hello-world/internal/handler"
	"github.com/desulaidovich/hello-world/internal/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Handler = fx.Module(
	"handler",
	fx.Provide(http.NewServeMux),
	fx.Provide(func(mux *http.ServeMux) http.Handler {
		return middleware.Chain(mux, middleware.RequestID())
	}),
	fx.Provide(handler.NewHealth),
	fx.Invoke(func(root *http.ServeMux, health *handler.HealthHandler, log *zap.Logger) {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /health", health.Say)
		root.Handle("/health", middleware.Chain(
			mux,
			middleware.Logging(log),
		))
	}),
)
