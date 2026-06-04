package config

import v "github.com/go-ozzo/ozzo-validation/v4"

// Path — путь к файлу конфигурации на файловой системе.
type Path string

type Config struct {
	Development bool `koanf:"development"`
	HTTP        HTTP `koanf:"http"`
}

func (c Config) Validate() error {
	return v.ValidateStruct(
		&c,
		v.Field(&c.HTTP),
	)
}
