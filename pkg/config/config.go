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
	BucketTypeName
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

func (c Config) BucketName(dataType BucketTypeName) string {
	switch dataType {
	case BucketTypeSpec:
		return c.Spec.BucketName()
	case BucketTypeSot:
		return c.Sot.BucketName()
	default:
		return c.Sor.BucketName()
	}
}

func (c Config) Query(dataType BucketTypeName) string {
	switch dataType {
	case BucketTypeSpec:
		return c.Spec.Query()
	case BucketTypeSot:
		return c.Sot.Query()
	default:
		return c.Sor.Query()
	}
}

func (c Config) File(dataType BucketTypeName) string {
	switch dataType {
	case BucketTypeSpec:
		return c.Spec.File()
	case BucketTypeSot:
		return c.Sot.File()
	default:
		return c.Sor.File()
	}
}

func (c Config) Active(dataType BucketTypeName) bool {
	switch dataType {
	case BucketTypeSpec:
		return c.Spec.Active()
	case BucketTypeSot:
		return c.Sot.Active()
	default:
		return true
	}
}
