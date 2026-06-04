package module

import (
	"context"
	"fmt"

	"github.com/desulaidovich/hello-world/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = fx.Module(
	"logger",
	fx.Provide(func(cfg *config.Config) (*zap.Logger, error) {
		var zapCfg zap.Config

		if cfg.Development {
			zapCfg = zap.NewDevelopmentConfig()
			zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
			zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.000")
			zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			zapCfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		} else {
			zapCfg = zap.NewProductionConfig()
			zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
			zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
			zapCfg.EncoderConfig.TimeKey = "timestamp"
		}

		log, err := zapCfg.Build()
		if err != nil {
			return nil, fmt.Errorf("logger: build: %w", err)
		}
		return log, nil
	}),
	fx.Invoke(func(lc fx.Lifecycle, log *zap.Logger) {
		lc.Append(fx.Hook{
			OnStop: func(_ context.Context) error {
				_ = log.Sync()
				return nil
			},
		})
	}),
)
