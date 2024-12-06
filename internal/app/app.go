package app

import "github.com/mizmorr/wallet/internal/router"

func Run() error {
	router.Handle()
	return nil
}
