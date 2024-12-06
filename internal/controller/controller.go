package controller

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ToPostgres() error {
	connStr := "postgres://postgres:post@db:5432/pgbounce"

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}

	fmt.Println(greeting)
	return nil
}

func ViaBouncer() error {
	connStr := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(greeting)
	return nil
}
