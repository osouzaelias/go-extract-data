package config

import (
	"os"
)

type Config struct {
	*Sor
	*Sot
	*Spec
	*Table
	Region string
}

func NewConfig() *Config {
	return &Config{
		Sor:    NewSor(),
		Sot:    NewSot(),
		Spec:   NewSpec(),
		Table:  NewTable(),
		Region: os.Getenv("REGION"),
	}
}
