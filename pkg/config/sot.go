package config

import (
	"os"
	"strings"
)

type Sot struct {
	bucketName string
	file       string
	query      string
	active     bool
}

func NewSot() *Sot {
	return &Sot{
		bucketName: os.Getenv("SOT_NAME"),
		file:       os.Getenv("SOT_FILE"),
		query:      os.Getenv("SOT_QUERY"),
		active:     len(strings.TrimSpace(os.Getenv("SOT_NAME"))) > 0,
	}
}

func (s Sot) BucketName() string {
	return s.bucketName
}

func (s Sot) File() string {
	return s.file
}

func (s Sot) Query() string {
	return s.query
}

func (s Sot) Active() bool {
	return s.active
}
