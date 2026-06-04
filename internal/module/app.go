package module

import (
	"go.uber.org/fx"
)

var HTTPServer = fx.Options(
	Config,
	Logger,
	Handler,
	Server,
)
