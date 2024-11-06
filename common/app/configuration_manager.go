package app

import (
	"product-app/common/postgresql"
	"time"
)

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgresConfig := GetPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgresConfig,
	}
}

func GetPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "127.0.0.1",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: time.Second * 30,
	}
}
