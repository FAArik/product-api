package postgresql

import "time"

type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DbName                string
	MaxConnections        string
	MaxConnectionIdleTime time.Duration
}
