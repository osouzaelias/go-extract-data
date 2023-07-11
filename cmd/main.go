package main

import (
	"go-extract-data/internal"
	"go-extract-data/pkg/config"
	"log"
)

func main() {
	cfg := config.NewConfig()
	extract := internal.NewExtractData(*cfg)

	go func() {
		if err := extract.Extract(internal.Sor); err != nil {
			log.Println("Error > Extract SOR >", err)
		}
	}()

	if err := extract.Extract(internal.Sor); err != nil {
		log.Println("Error > Extract SOR >", err)
	}

	log.Println("Execução concluída com sucesso.")
}
