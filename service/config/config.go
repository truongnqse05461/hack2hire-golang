package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenPort    string `envconfig:"HTTP_PORT"`
	MysqlURL      string `ignore:"true"`
	MysqlHost     string `envconfig:"MYSQL_HOST"`
	MysqlDatabase string `envconfig:"MYSQL_DATABASE"`
	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
}

func GetConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	config.buildMysqlURL()
	return &config, nil
}

func (c *Config) buildMysqlURL() {
	c.MysqlURL = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.MysqlUser, c.MysqlPassword, c.MysqlHost, c.MysqlDatabase)
}
