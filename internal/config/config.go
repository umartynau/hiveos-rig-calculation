package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	DefaultConfigFile = "config"
)

type Config struct {
	App  App  `yaml:"app"`
	Hive Hive `yaml:"hive"`
}

type App struct {
	LogLevel        int       `yaml:"log_level"`
	LogColorEnabled bool      `yaml:"log_color_enabled"`
	LogTraceEnabled bool      `yaml:"log_trace_enabled"`
	InstanceName    string    `yaml:"instance_name"`
	InstanceLabel   string    `yaml:"instance_label"`
	Telemetry       Telemetry `yaml:"telemetry"`
}

type Telemetry struct {
	Enabled      bool   `yaml:"enabled"`
	NewrelicName string `yaml:"nr_name"`
	NewrelicKey  string `yaml:"nr_key"`
	DebugEnabled bool   `yaml:"debug_enabled"`
}

type Hive struct {
	ApiUrl string `yaml:"api_url"`
	ApiKey string `yaml:"api_key"`
}

func NewConfig(filename string) (*Config, error) {
	if len(filename) == 0 {
		filename = DefaultConfigFile
	}
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Panic().Err(err).Str("filename", filename).Msg("failed to read config file")
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		log.Panic().Err(err).Str("filename", filename).Msg("failed to unmarshal yaml")
		return nil, err
	}

	return cfg, err
}
