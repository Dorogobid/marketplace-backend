package config

import (
	"fmt"
	"net/url"
)

type DBConfig struct {
	Username string
	Password string
	DBName   string
	DBPort   string
	DBHost   string
}

func (c *DBConfig) DSN() string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.Username, c.Password),
		Host:     fmt.Sprintf("%s:%s", c.DBHost, c.DBPort),
		Path:     c.DBName,
		RawQuery: "sslmode=disable",
	}
	return u.String()
}
