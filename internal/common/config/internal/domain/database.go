package domain

import (
	"fmt"
	"time"
)

type database struct {
	Host     string
	Port     int16
	User     string
	Password string
	DB       string
	MaxIdle  int `toml:"max_idle"`
	MaxOpen  int `toml:"max_open"`
}

func (db *database) GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4", db.User, db.Password, db.Host, db.Port, db.DB)
}
func (db *database) GetMaxIdle() int {
	return db.MaxIdle
}
func (db *database) GetMaxOpen() int {
	return db.MaxOpen
}
func (db *database) GetMaxConnectionLifetime() time.Duration {
	return time.Minute * 30
}
