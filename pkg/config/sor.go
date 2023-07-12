package config

import "os"

type Sor struct {
	bucketName string
	file       string
	query      string
}

func NewSor() *Sor {
	return &Sor{
		bucketName: os.Getenv("SOR_NAME"),
		file:       os.Getenv("SOR_FILE"),
		query:      os.Getenv("SOR_QUERY"),
	}
}

func (s Sor) BucketName() string {
	return s.bucketName
}

func (s Sor) File() string {
	return s.file
}

func (s Sor) Query() string {
	return s.query
}
