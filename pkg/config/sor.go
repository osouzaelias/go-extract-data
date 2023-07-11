package config

import "os"

type Sor struct {
	name  string
	file  string
	query string
}

func NewSor() *Sor {
	return &Sor{
		name:  os.Getenv("SOR_NAME"),
		file:  os.Getenv("SOR_FILE"),
		query: os.Getenv("SOR_QUERY"),
	}
}

func (s Sor) Name() string {
	return s.name
}

func (s Sor) File() string {
	return s.file
}

func (s Sor) Query() string {
	return s.query
}
