package module

import (
	"github.com/desulaidovich/hello-world/config"
	"go.uber.org/fx"
)

var Config = fx.Module(
	"config",
	fx.Provide(config.NewProvider),
)
