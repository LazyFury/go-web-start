package config

import (
	"fmt"
)

type mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func (m *mysql) ToString() string {
	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True"
	return fmt.Sprintf(format, m.User, m.Password, m.Host, m.Port, m.Database) + "&loc=Asia%2FShanghai"
}
