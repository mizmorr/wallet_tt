package main

import (
	"github.com/mizmorr/wallet/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
