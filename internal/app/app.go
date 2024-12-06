package app

import (
	"log"

	"github.com/mizmorr/wallet/config"
	"github.com/mizmorr/wallet/internal/router"
)

func Run() error {
	cfg := config.Get()
	log.Println(cfg.DatabaseURL)
	router.Handle()

	return nil
}
