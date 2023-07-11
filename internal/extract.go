package internal

import (
	"encoding/json"
	"fmt"
	"go-extract-data/pkg/aws"
	"go-extract-data/pkg/config"
	"log"
	"os"
)

type DataType uint8

const (
	Sor DataType = iota
	Sot
	Spec
)

type ExtractData struct {
	storage  *aws.S3Client
	database *aws.DynamoDBClient
	config   config.Config
}

func NewExtractData(c config.Config) *ExtractData {
	return &ExtractData{
		storage:  aws.NewS3Client(c),
		database: aws.NewDynamoDBClient(c),
		config:   c,
	}
}

func (d ExtractData) Extract(dataType DataType) error {
	results, err := d.database.ExecuteStatement(d.getQuery(dataType))

	if err != nil {
		return fmt.Errorf("falha na query do DynamoDB: %s", err)
	}

	tempFile, err := os.CreateTemp(os.TempDir(), "prefix-")

	if err != nil {
		log.Fatalf("Falha ao criar arquivo temporário: %s", err)
	}

	defer os.Remove(tempFile.Name())

	encoder := json.NewEncoder(tempFile)

	for _, result := range results {
		err = encoder.Encode(&result)

		if err != nil {
			return fmt.Errorf("falha ao escrever no arquivo: %s", err)
		}
	}

	file, err := os.Open(tempFile.Name())

	if err != nil {
		return fmt.Errorf("falha ao abrir o arquivo: %s", err)
	}

	err = d.storage.PutObject(d.getBucketName(dataType), "resultado.json", file)

	if err != nil {
		return fmt.Errorf("falha ao enviar objeto para S3: %s", err)
	}

	return nil
}

func (d ExtractData) getBucketName(dataType DataType) string {
	switch dataType {
	case Spec:
		return d.config.Spec.Name()
	case Sot:
		return d.config.Sot.Name()
	default:
		return d.config.Sor.Name()
	}
}

func (d ExtractData) getQuery(dataType DataType) string {
	switch dataType {
	case Spec:
		return d.config.Spec.Query()
	case Sot:
		return d.config.Sot.Query()
	default:
		return d.config.Sor.Query()
	}
}
