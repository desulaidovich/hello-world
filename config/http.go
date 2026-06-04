package config

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type HTTP struct {
	Host              string        `koanf:"host"`
	Port              string        `koanf:"port"`
	ReadHeaderTimeout time.Duration `koanf:"read_header_timeout"`
	ReadTimeout       time.Duration `koanf:"read_timeout"`
	WriteTimeout      time.Duration `koanf:"write_timeout"`
	IdleTimeout       time.Duration `koanf:"idle_timeout"`
}

func (h HTTP) Validate() error {
	return v.ValidateStruct(
		&h,
		v.Field(&h.Port, v.Required),
		v.Field(&h.ReadHeaderTimeout, v.Required, v.Min(time.Duration(1))),
		v.Field(&h.ReadTimeout, v.Required, v.Min(time.Duration(1))),
		v.Field(&h.WriteTimeout, v.Required, v.Min(time.Duration(1))),
		v.Field(&h.IdleTimeout, v.Required, v.Min(time.Duration(1))),
	)
}
