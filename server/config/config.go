package config

import (
	"github.com/jinzhu/configor"
)

var (
	Config             ServiceConfig
	DefaultConfigFiles = []string{
		"config/database.yml",
		"config/smtp.yml",
		"config/application.yml",
		"config/cloud.yml",
		"config/plugins.yml",
		"config/tasks.yml",
	}
)

type ServiceConfig struct {
	Port  uint `env:"APP_PORT" default:"7000"`
	Debug bool `env:"APP_DEBUG" default:"false"`
	DB    struct {
		Name          string `env:"DB_NAME" default:"admin_on_rest_entities"`
		Adapter       string `env:"DB_ADAPTER" default:"sqlite"`
		Host          string `env:"DB_HOST" default:"localhost"`
		Port          string `env:"DB_PORT" default:"3306"`
		User          string `env:"DB_USER"`
		Password      string `env:"DB_PASSWORD"`
		PrefixPath    string `env:"DB_PREFIX_PATH" default:"./"`
		MigrationsDir string `env:"DB_MIGRATION_DIR" default:"./shared/data/seeds"`
		Secured       bool   `env:"DB_SECURED" default:"false"`
		SingularTable bool   `env:"DB_SINGULAR_TABLE" default:"false"`
		Debug         bool   `env:"DB_DEBUG" default:"false"`
		// DSN           string `env:"DB_DSN" default:"host=localhost user=ireflect password=1Reflect dbname=ireflect-dev sslmode=disable"`
	}
}

func init() {
	if err := configor.Load(&Config, DefaultConfigFiles...); err != nil {
		panic(err)
	}
}
