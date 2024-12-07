package controller

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mizmorr/wallet/config"
)

func ToPostgres() error {
	ctx := context.Background()
	confg := config.Get()
	poolConfig, err := pgxpool.ParseConfig(confg.DatabaseURL)

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	var greeting string
	err = db.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}

	fmt.Println(greeting)
	return nil
}
