package config

import (
	"github.com/jinzhu/configor"
	// "github.com/roscopecoltran/admin-on-rest-server/server/config/gateway"
	// "github.com/roscopecoltran/admin-on-rest-server/server/config/swagger"
	// "github.com/roscopecoltran/admin-on-rest-server/server/config/aws"
	// "github.com/roscopecoltran/admin-on-rest-server/server/config/smtp"
)

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `env:"DBName" default:"admin_on_rest_entities"`
		Adapter  string `env:"DBAdapter" default:"sqlite"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
	// S3   S3Config
	// SMTP SMTPConfig
}{}

var (
	DefaultConfigFiles = []string{"config/database.yml", "config/smtp.yml", "config/application.yml", "config/cloud.yml", "config/plugins.yml"}
)

func init() {
	if err := configor.Load(&Config, DefaultConfigFiles...); err != nil {
		panic(err)
	}
}
