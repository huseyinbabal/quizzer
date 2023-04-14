package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	DB  DB
	App App
}

type DB struct {
	Host     string
	User     string
	Password string
	Port     int64
	Name     string
	SslMode  string `yaml:"sslmode"`
}

func (d *DB) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", d.Host, d.User, d.Password, d.Name, d.Port, d.SslMode)
}

type App struct {
	Port int64
}

func Init() (*Config, error) {
	k := koanf.New(".")
	parser := yaml.Parser()

	configFile := "config.yml"
	if os.Getenv("CONFIG_LOCATION") != "" {
		configFile = os.Getenv("CONFIG_LOCATION")
	}
	if err := k.Load(file.Provider(configFile), parser); err != nil {
		return nil, errors.Wrap(err, "while loading config file.")
	}

	var cfg Config
	err := k.Unmarshal("", &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "while unmarshalling config file.")
	}
	return &cfg, nil
}
