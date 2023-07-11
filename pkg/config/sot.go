package config

import (
	"os"
	"strings"
)

type Sot struct {
	name   string
	file   string
	query  string
	active bool
}

func NewSot() *Sot {
	return &Sot{
		name:   os.Getenv("SOT_NAME"),
		file:   os.Getenv("SOT_FILE"),
		query:  os.Getenv("SOT_QUERY"),
		active: len(strings.TrimSpace(os.Getenv("SOT_NAME"))) > 0,
	}
}

func (s Sot) Name() string {
	return s.name
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
