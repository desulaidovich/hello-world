package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func NewProvider(yamlPath Path) (*Config, error) {
	if yamlPath == "" {
		return nil, fmt.Errorf("config: yamlPath is required")
	}

	k := koanf.New(".")

	if err := k.Load(file.Provider(string(yamlPath)), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("config: load yaml: %w", err)
	}

	// Двойное подчёркивание — разделитель пути, одиночное допустимо в именах полей:
	// HTTP__READ_HEADER_TIMEOUT -> http.read_header_timeout
	// DATABASE__USER__NAME -> database.user.name
	for _, prefix := range []string{"HTTP__", "DATABASE__"} {
		if err := k.Load(env.Provider(prefix, ".", func(key string) string {
			return strings.ToLower(strings.ReplaceAll(key, "__", "."))
		}), nil); err != nil {
			return nil, fmt.Errorf("config: load env (%s): %w", prefix, err)
		}
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("config: unmarshal: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config: validate: %w", err)
	}

	return &cfg, nil
}
