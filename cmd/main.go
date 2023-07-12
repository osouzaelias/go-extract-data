package main

import (
	"go-extract-data/internal"
	"go-extract-data/pkg/config"
	"log"
	"sync"
)

func main() {
	cfg := config.NewConfig()
	extract := internal.NewExtractData(*cfg)

	var wg sync.WaitGroup
	for _, name := range cfg.BucketTypeNames() {
		if cfg.Active(name) {
			wg.Add(1)
			go executeExtract(extract, name, &wg)
		}
	}

	wg.Wait()

	log.Println("Execução concluída com sucesso.")
}

func executeExtract(extract *internal.ExtractData, dataType config.BucketTypeName, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := extract.Extract(dataType); err != nil {
		log.Printf("Error > Extract > %s > %v", dataType, err)
	}
}
