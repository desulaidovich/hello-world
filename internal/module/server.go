package module

import (
	"context"
	"fmt"

	"github.com/desulaidovich/hello-world/internal/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type serverParams struct {
	fx.In

	Lifecycle  fx.Lifecycle
	Server     *server.Server
	Log        *zap.Logger
	Shutdowner fx.Shutdowner
}

var Server = fx.Module(
	"http server",
	fx.Provide(server.New),
	fx.Invoke(func(p serverParams) {
		p.Lifecycle.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				p.Log.Info("starting http server")
				go func() {
					if err := p.Server.Start(); err != nil {
						p.Log.Error("http server failed", zap.Error(err))
						_ = p.Shutdowner.Shutdown(fx.ExitCode(1))
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				p.Log.Info("stopping http server")
				if err := p.Server.Stop(ctx); err != nil {
					return fmt.Errorf("http server: %w", err)
				}
				return nil
			},
		})
	}),
)
