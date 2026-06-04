// Точка входа HTTP-сервера.
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/desulaidovich/hello-world/config"
	"github.com/desulaidovich/hello-world/internal/module"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "HTTP server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.yaml",
				Usage:   "path to config file",
			},
		},
		Action: func(c *cli.Context) error {
			return run(c.String("config"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "run: %v\n", err)
		os.Exit(1)
	}
}

func run(yamlPath string) error {
	app := fx.New(
		fx.Supply(config.Path(yamlPath)),
		module.HTTPServer,
		fx.WithLogger(func(log *zap.Logger, cfg *config.Config) fxevent.Logger {
			if cfg.Development {
				return &fxevent.ZapLogger{
					Logger: log.Named("fx"),
				}
			}
			return fxevent.NopLogger
		}),
	)

	app.Run()

	if err := app.Err(); err != nil {
		return fmt.Errorf("app: %w", err)
	}
	return nil
}
