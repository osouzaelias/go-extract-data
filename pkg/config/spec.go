package config

import (
	"os"
	"strings"
)

type Spec struct {
	name   string
	file   string
	query  string
	active bool
}

func NewSpec() *Spec {
	return &Spec{
		name:   os.Getenv("SPEC_NAME"),
		file:   os.Getenv("SPEC_FILE"),
		query:  os.Getenv("SPEC_QUERY"),
		active: len(strings.TrimSpace(os.Getenv("SPEC_NAME"))) > 0,
	}
}

func (s Spec) Name() string {
	return s.name
}

func (s Spec) File() string {
	return s.file
}

func (s Spec) Query() string {
	return s.query
}

func (s Spec) Active() bool {
	return s.active
}
