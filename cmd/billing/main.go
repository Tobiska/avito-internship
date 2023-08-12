package main

import (
	"github.com/Tobiska/avito-internship/config"
	"github.com/Tobiska/avito-internship/internal/app/billing"
	"log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	billing.Run(cfg)
}
