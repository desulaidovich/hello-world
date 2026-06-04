package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/desulaidovich/hello-world/config"
	"go.uber.org/zap"
)

type Server struct {
	httpSrv *http.Server
	logger  *zap.Logger
}

func New(logger *zap.Logger, handler http.Handler, cfg *config.Config) *Server {
	return &Server{
		logger: logger,
		httpSrv: &http.Server{
			Addr:              net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port),
			Handler:           handler,
			ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
			ReadTimeout:       cfg.HTTP.ReadTimeout,
			WriteTimeout:      cfg.HTTP.WriteTimeout,
			IdleTimeout:       cfg.HTTP.IdleTimeout,
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info(
		"listening",
		zap.String("addr", s.httpSrv.Addr),
	)

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
